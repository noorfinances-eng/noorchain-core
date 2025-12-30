package node

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/syndtr/goleveldb/leveldb"

	"encoding/hex"
	"encoding/json"
	"noorchain-evm-l1/core/config"
	"noorchain-evm-l1/core/evmstate"
	"noorchain-evm-l1/core/exec"
	"noorchain-evm-l1/core/health"
	"noorchain-evm-l1/core/network"
	"noorchain-evm-l1/core/txindex"
	"noorchain-evm-l1/core/txpool"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/holiman/uint256"
	"github.com/ethereum/go-ethereum/core/tracing"

)

type Logger interface {
	Println(v ...any)
}

type State string

const (
	StateInit     State = "INIT"
	StateRunning  State = "RUNNING"
	StateStopping State = "STOPPING"
)

// ---- M9 receipts persistence (minimal) ----

const rcptKVPrefix = "rcpt/v1/"

// EVM Chain ID (EIP-155) must be tooling/wallet compatible.
// For NOORCHAIN 2.1 local dev, we pin this to 2121.
const evmChainID uint64 = 2121

const stateHeadKVKey = "stateroot/v1/head" // persisted in NOOR LevelDB (n.db)

func rcptKey(txHash string) []byte {
	h := strings.ToLower(strings.TrimPrefix(strings.TrimSpace(txHash), "0x"))
	return []byte(rcptKVPrefix + h)
}

type receiptJSON struct {
	TransactionHash   common.Hash     `json:"transactionHash"`
	TransactionIndex  string          `json:"transactionIndex"`
	BlockHash         common.Hash     `json:"blockHash"`
	BlockNumber       string          `json:"blockNumber"`
	From              common.Address  `json:"from"`
	To                *common.Address `json:"to"`
	CumulativeGasUsed string          `json:"cumulativeGasUsed"`
	GasUsed           string          `json:"gasUsed"`
	ContractAddress   *common.Address `json:"contractAddress"`
	Logs              []any           `json:"logs"`
	Status            string          `json:"status"`
	Type              string          `json:"type"`
}

func pseudoBlockHash(n uint64) common.Hash {
	b := make([]byte, 32)
	for i := 0; i < 8; i++ {
		b[31-i] = byte(n >> (8 * i))
	}
	return common.BytesToHash(crypto.Keccak256(b))
}

func toHexBig(v uint64) string {
	// v is uint64 (gas), represent as hex quantity
	return "0x" + fmt.Sprintf("%x", v)
}

func toHexUint(v uint64) string {
	return "0x" + fmt.Sprintf("%x", v)
}

type Node struct {
	cfg     config.Config
	logger  Logger
	network *network.Network
	health  *health.Server

	txpool  *txpool.Pool
	txindex *txindex.Index

	db *leveldb.DB

	// M12: geth-compatible DB for Ethereum world-state (trie/code/storage), isolated under <data-dir>/db/geth
	evmStore *evmstate.Store

	ctx    context.Context
	cancel context.CancelFunc

	mu     sync.Mutex
	state  State
	height uint64
}

func (n *Node) DB() *leveldb.DB { return n.db }

func New(cfg config.Config) *Node {
	ctx, cancel := context.WithCancel(context.Background())
	return &Node{
		cfg:     cfg,
		logger:  newLogger(),
		network: network.New(cfg.P2PAddr),
		health:  health.New(cfg.HealthAddr),
		txpool:  txpool.New(),
		txindex: txindex.New(),
		ctx:     ctx,
		cancel:  cancel,
		state:   StateInit,
		height:  0,
	}
}

func (n *Node) Start() error {
	if err := os.MkdirAll(n.cfg.DataDir, 0o755); err != nil {
		return fmt.Errorf("mkdir data-dir %s: %w", n.cfg.DataDir, err)
	}

	db, err := openLevelDB(n.cfg.DataDir)
	if err != nil {
		return err
	}
	n.db = db

	// M12: open isolated geth DB for EVM world-state (Option A)
	es, err := evmstate.Open(n.cfg.DataDir, false)
	if err != nil {
		_ = n.db.Close()
		n.db = nil
		return err
	}
	n.evmStore = es

	n.mu.Lock()
	n.state = StateRunning
	n.mu.Unlock()

	n.logger.Println("node started")
	n.logger.Println("state:", StateRunning)
	n.logger.Println("chain-id:", n.cfg.ChainID)
	n.logger.Println("data-dir:", n.cfg.DataDir)

	if err := n.network.Start(); err != nil {
		return err
	}

	// M10: dial boot peers (best-effort)
	for _, peer := range n.cfg.BootPeers {
		_ = n.network.Connect(peer)
	}
	if err := n.health.Start(); err != nil {
		return err
	}

	go n.loop()
	return nil
}

func (n *Node) loop() {
	// M10: follower mode (mainnet-like pack): follow leader height via RPC, do not mine locally
	if strings.EqualFold(strings.TrimSpace(n.cfg.Role), "follower") && strings.TrimSpace(n.cfg.FollowRPC) != "" {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		type rpcResp struct {
			Result string `json:"result"`
		}

		for {
			select {
			case <-n.ctx.Done():
				n.logger.Println("node loop stopped")
				return
			case <-ticker.C:
				reqBody := `{"jsonrpc":"2.0","id":1,"method":"eth_blockNumber","params":[]}`
				resp, err := http.Post(strings.TrimRight(n.cfg.FollowRPC, "/"), "application/json", strings.NewReader(reqBody))
				if err != nil {
					n.logger.Println("follower: follow-rpc post error:", err)
					continue
				}
				b, _ := io.ReadAll(resp.Body)
				_ = resp.Body.Close()

				var rr rpcResp
				if err := json.Unmarshal(b, &rr); err != nil {
					n.logger.Println("follower: decode error:", err)
					continue
				}
				// parse hex quantity
				h := strings.TrimSpace(rr.Result)
				h = strings.TrimPrefix(h, "0x")
				if h == "" {
					continue
				}
				hv, err := strconv.ParseUint(h, 16, 64)
				if err != nil {
					n.logger.Println("follower: parse height error:", err)
					continue
				}

				n.mu.Lock()
				if hv > n.height {
					n.height = hv
				}
				height := n.height
				nodeState := n.state
				n.mu.Unlock()

				n.logger.Println("tick | height:", height, "| state:", nodeState, "| follower:", true)
			}
		}
	}

	// M12.2: init geth world-state (StateDB + triedb) using isolated geth DB.
	headRoot := common.Hash{}
	if n.db != nil {
		if b, err := n.db.Get([]byte(stateHeadKVKey), nil); err == nil && len(b) == 32 {
			headRoot = common.BytesToHash(b)
		}
	}

	var tdb *triedb.Database
	var sdbCache *state.CachingDB
	var statedb *state.StateDB

	if n.evmStore != nil && n.evmStore.DB() != nil {
		diskdb := rawdb.NewDatabase(n.evmStore.DB())
			tdb = triedb.NewDatabase(diskdb, nil)
		sdbCache = state.NewDatabase(tdb, nil)
		if st, err := state.New(headRoot, sdbCache); err == nil {
			statedb = st
		} else {
			n.logger.Println("evmstate: state.New failed, fallback empty root | err:", err)
			headRoot = common.Hash{}
			statedb, _ = state.New(headRoot, sdbCache)
		}
	} else {
		n.logger.Println("evmstate: missing evmStore DB; stateRoot will remain placeholder")
	}

	// M9: execution hook (minimal, step 1)
	// For now: decode raw tx bytes and log decoded shape.
	// Next steps will route contract calls + build receipts + persist.
	applyTx := func(t txpool.Tx, height uint64) *types.Receipt {
		if len(t.Raw) == 0 {
			n.logger.Println("applyTx: empty raw tx | hash:", t.Hash, "| height:", height)
			return nil
		}

		var tx types.Transaction
		if err := tx.UnmarshalBinary(t.Raw); err != nil {
			n.logger.Println("applyTx: decode failed | hash:", t.Hash, "| height:", height, "| err:", err)
			return nil
		}

		// Best-effort "to" for logging (contract creation => nil)
		toStr := "<create>"
		if tx.To() != nil {
			toStr = tx.To().Hex()
		}

		// Sanity: ensure pool hash matches raw bytes hash
		rawHash := crypto.Keccak256Hash(t.Raw)
		match := strings.EqualFold(rawHash.Hex(), t.Hash)

		// Sender is required for receipt.from and contractAddress derivation (CREATE).
		from := common.Address{}
		signer := types.LatestSignerForChainID(new(big.Int).SetUint64(evmChainID))
		if f, err := types.Sender(signer, &tx); err == nil {
			from = f
		} else {
			n.logger.Println("applyTx: sender decode failed | hash:", t.Hash, "| err:", err)
		}
		// M12.2: minimal world-state write (nonce) so RPC can read real values.
		// Not full EVM execution; just proves StateDB is live and persisted.
		if statedb != nil && from != (common.Address{}) {
		statedb.SetNonce(from, tx.Nonce()+1)
		
		// M12.2: minimal balance write for proof (temporary until full value-transfer/EVM).
		statedb.AddBalance(from, uint256.NewInt(1), tracing.BalanceChangeUnspecified)
		}

		// Contract creation => contractAddress = CreateAddress(from, nonce)
		var contractAddrPtr *common.Address
		if tx.To() == nil && from != (common.Address{}) {
			ca := crypto.CreateAddress(from, tx.Nonce())
			contractAddrPtr = &ca
		}

		// M9: persist minimal receipt (best-effort)
		if n.db != nil {
			toPtr := tx.To()
			rcpt := receiptJSON{
				TransactionHash:   rawHash,
				TransactionIndex:  "0x0",
				BlockHash:         pseudoBlockHash(height),
				BlockNumber:       toHexUint(height),
				From:              from,
				To:                toPtr,
				CumulativeGasUsed: toHexBig(tx.Gas()),
				GasUsed:           toHexBig(tx.Gas()),
				ContractAddress:   contractAddrPtr,
				Logs:              []any{},
				Status:            "0x1",
				Type:              toHexUint(uint64(tx.Type())),
			}
			if b, err := json.Marshal(rcpt); err == nil {
				_ = n.db.Put(rcptKey(rawHash.Hex()), b, nil)
			}
		}

		// M9: minimal contracts execution (PoSS submitSnapshot)
		possOK := false
		if n.db != nil {
			if ok, err := exec.ApplyPoSSSubmitSnapshot(&tx, n.cfg.ChainID, n.db, uint64(time.Now().Unix())); err != nil {
				n.logger.Println("applyTx: poss exec error:", err)
			} else {
				possOK = ok
			}
		}

		// M12: build a minimal geth receipt for per-block receiptsRoot/logsBloom
		rcptGeth := &types.Receipt{
			TxHash:            rawHash,
			BlockHash:         pseudoBlockHash(height),
			BlockNumber:       new(big.Int).SetUint64(height),
			TransactionIndex:  0,
			Status:            1,
			CumulativeGasUsed: tx.Gas(),
			GasUsed:           tx.Gas(),
			Logs:              []*types.Log{},
		}

		n.logger.Println(
			"applyTx: DECODE_OK",
			"| height:", height,
			"| poolHash:", t.Hash,
			"| rawHash:", rawHash.Hex(),
			"| match:", match,
			"| type:", tx.Type(),
			"| nonce:", tx.Nonce(),
			"| to:", toStr,
			"| dataLen:", len(tx.Data()),
			"| possOK:", possOK,
		)
		return rcptGeth
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-n.ctx.Done():
			n.logger.Println("node loop stopped")
			return
		case <-ticker.C:
			n.mu.Lock()
			n.height++
			nodeState := n.state
			height := n.height
			n.mu.Unlock()

			// M8.A minimal inclusion: mark some pending txs as mined at this height
			mined := 0
			receipts := types.Receipts{}
			if n.txpool != nil && n.txindex != nil {
				txs := n.txpool.PopPending(64)
				for i := range txs {
					n.txindex.Put(txs[i].Hash, height)
					// M9: hook point (execution will be extended in next steps)
					r := applyTx(txs[i], height)
					if r != nil {
						receipts = append(receipts, r)
					}
					mined++
				}
			}

			// M12: persist block metadata (roots/bloom) for RPC reads
			if n.db != nil {
				bloom := types.CreateBloom(receipts)
				receiptsRoot := types.DeriveSha(receipts, trie.NewStackTrie(nil))

				// M12.2: compute stateRoot from geth StateDB + triedb (persisted), fallback to placeholder
				stateRoot := pseudoBlockHash(height)
				if statedb != nil {
					root, err := statedb.Commit(height, false)
					if err != nil {
						n.logger.Println("evmstate: commit failed, keep placeholder | err:", err)
					} else {
						stateRoot = root
						if tdb != nil {
							_ = tdb.Commit(root, false)
						}
						_ = n.db.Put([]byte(stateHeadKVKey), root.Bytes(), nil)
						if sdbCache != nil {
							if st2, err := state.New(root, sdbCache); err == nil {
								statedb = st2
							} else {
								n.logger.Println("evmstate: reopen state failed | err:", err)
							}
						}
					}
				}

				bm := blockMeta{
					Height:       height,
					BlockHash:    pseudoBlockHash(height),
					StateRoot:    stateRoot,
					ReceiptsRoot: receiptsRoot,
					LogsBloomHex: "0x" + hex.EncodeToString(bloom[:]),
				}
				if b, err := encodeBlockMeta(bm); err == nil {
					_ = n.db.Put(blkMetaKey(height), b, nil)
				}
			}

			if mined > 0 {
				n.logger.Println("tick | height:", height, "| state:", nodeState, "| mined:", mined)
			} else {
				n.logger.Println("tick | height:", height, "| state:", nodeState)
			}
		}
	}
}

func (n *Node) Config() config.Config { return n.cfg }

func (n *Node) Height() uint64 {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.height
}

func (n *Node) TxPool() *txpool.Pool    { return n.txpool }
func (n *Node) TxIndex() *txindex.Index { return n.txindex }
func (n *Node) EVMStore() *evmstate.Store { return n.evmStore }

// StateRootHead returns the latest committed world-state root (or zero-hash if missing).
func (n *Node) StateRootHead() common.Hash {
	if n == nil || n.db == nil {
		return common.Hash{}
	}
	b, err := n.db.Get([]byte(stateHeadKVKey), nil)
	if err != nil || len(b) != 32 {
		return common.Hash{}
	}
	return common.BytesToHash(b)
}


func (n *Node) Stop() error {
	n.mu.Lock()
	if n.state == StateStopping {
		n.mu.Unlock()
		return nil
	}
	n.state = StateStopping
	n.mu.Unlock()

	n.logger.Println("state:", StateStopping)

	n.cancel()

	// stop health first (fast)
	{
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_ = n.health.Stop(ctx)
	}

	// stop network
	if err := n.network.Stop(); err != nil {
		n.logger.Println("network stop error:", err)
	}

	if n.evmStore != nil {
		_ = n.evmStore.Close()
		n.evmStore = nil
	}

	if n.db != nil {
		_ = n.db.Close()
		n.db = nil
	}

	n.logger.Println("node stopped")
	return nil
}
