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
	"github.com/syndtr/goleveldb/leveldb/util"

	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"noorchain-evm-l1/core/config"
	"noorchain-evm-l1/core/evmstate"
	"noorchain-evm-l1/core/health"
	"noorchain-evm-l1/core/network"
	"noorchain-evm-l1/core/txindex"
	"noorchain-evm-l1/core/txpool"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/holiman/uint256"
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

const (
	rcptKVPrefix       = "rcpt/v1/"
	logRecKVPrefix     = "logrec/v1/"
	logRecAppliedKVKey = "logrec/v1/applied" // set after one-time backfill from receipts
)

// EVM Chain ID (EIP-155) must be tooling/wallet compatible.
// For NOORCHAIN 2.1 local dev, we pin this to 2121.
const evmChainID uint64 = 2121
const stateHeadKVKey = "stateroot/v1/head" // persisted in NOOR LevelDB (n.db)
const headHeightKVKey = "headheight/v1"    // persisted in NOOR LevelDB (n.db)

func u64be(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func readHeadHeight(db *leveldb.DB, headRoot common.Hash) (uint64, bool) {
	if db == nil {
		return 0, false
	}

	// 1) Preferred: direct key
	if b, err := db.Get([]byte(headHeightKVKey), nil); err == nil && len(b) == 8 {
		return binary.BigEndian.Uint64(b), true
	}

	// 2) Recovery: scan blkmeta and try to find the height whose StateRoot matches headRoot.
	// Also compute max height as fallback.
	var maxH uint64
	var matchH uint64
	var hasMatch bool

	it := db.NewIterator(util.BytesPrefix([]byte(blkMetaPrefix)), nil)
	defer it.Release()

	for it.Next() {
		k := string(it.Key()) // "blkmeta/v1/<hex>"
		suffix := strings.TrimPrefix(k, blkMetaPrefix)
		if suffix == "" {
			continue
		}
		hv, err := strconv.ParseUint(suffix, 16, 64)
		if err != nil {
			continue
		}
		if hv > maxH {
			maxH = hv
		}

		if headRoot != (common.Hash{}) {
			bm, err := decodeBlockMeta(it.Value())
			if err == nil && bm.StateRoot == headRoot {
				matchH = hv
				hasMatch = true
			}
		}
	}

	if hasMatch {
		_ = db.Put([]byte(headHeightKVKey), u64be(matchH), nil)
		return matchH, true
	}
	if maxH > 0 {
		_ = db.Put([]byte(headHeightKVKey), u64be(maxH), nil)
		return maxH, true
	}
	return 0, false
}

func rcptKey(txHash string) []byte {
	h := strings.ToLower(strings.TrimPrefix(strings.TrimSpace(txHash), "0x"))
	return []byte(rcptKVPrefix + h)
}
func logRecKey(height, txIndex, logIndex uint64) []byte {
	// Key ordering: heightBE | txIndexBE | logIndexBE
	// Lexicographic order == canonical order for iteration.
	prefix := []byte(logRecKVPrefix)
	b := make([]byte, len(prefix)+24)
	copy(b, prefix)
	binary.BigEndian.PutUint64(b[len(prefix)+0:], height)
	binary.BigEndian.PutUint64(b[len(prefix)+8:], txIndex)
	binary.BigEndian.PutUint64(b[len(prefix)+16:], logIndex)
	return b
}

func parseHexQtyString(s string) (uint64, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, false
	}
	s = strings.TrimPrefix(s, "0x")
	if s == "" {
		return 0, false
	}
	n, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return 0, false
	}
	return n, true
}

func parseHexQtyAny(v any) (uint64, bool) {
	s, ok := v.(string)
	if !ok {
		return 0, false
	}
	return parseHexQtyString(s)
}

// backfillLogRecIndex builds logrec/v1/* from persisted receipts (rcpt/v1/*) once per data-dir.
func backfillLogRecIndex(db *leveldb.DB, lg Logger) {
	if db == nil {
		return
	}
	if _, err := db.Get([]byte(logRecAppliedKVKey), nil); err == nil {
		return
	}

	it := db.NewIterator(util.BytesPrefix([]byte(rcptKVPrefix)), nil)
	defer it.Release()

	var total uint64
	for it.Next() {
		var r receiptJSON
		if err := json.Unmarshal(it.Value(), &r); err != nil {
			continue
		}

		bn, ok := parseHexQtyString(r.BlockNumber)
		if !ok {
			continue
		}

		for _, la := range r.Logs {
			lm, ok := la.(map[string]any)
			if !ok {
				continue
			}

			txi, ok := parseHexQtyAny(lm["transactionIndex"])
			if !ok {
				txi2, ok2 := parseHexQtyString(r.TransactionIndex)
				if !ok2 {
					continue
				}
				txi = txi2
			}
			li, ok := parseHexQtyAny(lm["logIndex"])
			if !ok {
				continue
			}

			if b, err := json.Marshal(lm); err == nil {
				_ = db.Put(logRecKey(bn, txi, li), b, nil)
				total++
			}
		}
	}

	_ = db.Put([]byte(logRecAppliedKVKey), []byte{1}, nil)
	if lg != nil {
		lg.Println("logrec: backfill applied | entries:", total)
	}
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
	// Restore chain head height from DB (or recover from blkmeta if missing)
	headRoot := common.Hash{}
	if b, err := n.db.Get([]byte(stateHeadKVKey), nil); err == nil && len(b) == 32 {
		headRoot = common.BytesToHash(b)
	}
	if hh, ok := readHeadHeight(n.db, headRoot); ok && hh > 0 {
		n.mu.Lock()
		n.height = hh
		n.mu.Unlock()
		n.logger.Println("head: restored | height:", hh, "| root:", headRoot.Hex())
	}

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

	// M12.x: restore head height from NOOR LevelDB (preferred key, else recover from blkmeta scan)
	if n.db != nil {
		if hh, ok := readHeadHeight(n.db, headRoot); ok {
			n.mu.Lock()
			n.height = hh
			n.mu.Unlock()

			// If headRoot was zero (or not yet loaded), refresh it from the persisted head root key.
			if headRoot == (common.Hash{}) {
				if b, err := n.db.Get([]byte(stateHeadKVKey), nil); err == nil && len(b) == 32 {
					headRoot = common.BytesToHash(b)
				}
			}

			n.logger.Println("head: restored | height:", hh, "| root:", headRoot.Hex())
		}
	}

	// M12.5.1: optional alloc bootstrap (dev/mainnet genesis allocations)
	// Applied once per data-dir (guarded by alloc/v1/applied in NOOR LevelDB).
	if statedb != nil && n.db != nil {
		if strings.TrimSpace(n.cfg.AllocFile) != "" {
			if _, err := n.db.Get([]byte(allocAppliedKVKey), nil); err != nil {
				af, err := readAllocFile(n.cfg.AllocFile)
				if err != nil {
					n.logger.Println("alloc: read failed:", err)
				} else if af.ChainID != evmChainID {
					n.logger.Println("alloc: chainId mismatch | file:", af.ChainID, "expected:", evmChainID)
				} else if allocs, err := parseAllocEntries(af); err != nil {
					n.logger.Println("alloc: parse failed:", err)
				} else {
					n.logger.Println("alloc: applying", len(allocs), "entries from", n.cfg.AllocFile)
					for addr, bal := range allocs {
						statedb.SetBalance(addr, bal, tracing.BalanceChangeUnspecified)
					}
					// Commit allocation and persist head root
					if root, err := statedb.Commit(0, false); err == nil {
						_ = tdb.Commit(root, false)
						_ = n.db.Put([]byte(stateHeadKVKey), root.Bytes(), nil)
						_ = n.db.Put([]byte(headHeightKVKey), u64be(0), nil)
						_ = n.db.Put([]byte(allocAppliedKVKey), []byte{1}, nil)
						if st2, err := state.New(root, sdbCache); err == nil {
							statedb = st2
						}
						n.logger.Println("alloc: applied | new head root:", root.Hex())
					} else {
						n.logger.Println("alloc: commit failed:", err)
					}

				}
			}
		}
	}
	// M14-B: one-time backfill of log index from persisted receipts (safe, idempotent).
	if n.db != nil {
		backfillLogRecIndex(n.db, n.logger)
	}

	// M9: execution hook (minimal, step 1)
	// For now: decode raw tx bytes and log decoded shape.
	// Next steps will route contract calls + build receipts + persist.
	applyTx := func(t txpool.Tx, height uint64, txIndex int, blockLogBase *uint64) *types.Receipt {
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
		// ---- M13: real EVM execution via geth state transition (mainnet-like) ----
		blockHash := pseudoBlockHash(height)

		// Must have statedb for real execution
		if statedb == nil || from == (common.Address{}) {
			n.logger.Println("applyTx: missing statedb/from | hash:", t.Hash, "| height:", height)
			return nil
		}

		// Set tx context so logs are indexed correctly inside StateDB
		statedb.SetTxContext(rawHash, txIndex)

		// Chain config: enable all forks from genesis (needed for SHL, etc.), but enforce our chainId
		cc := *params.AllDevChainProtocolChanges
		cc.ChainID = new(big.Int).SetUint64(evmChainID)

		num := new(big.Int).SetUint64(height)
		ts := uint64(time.Now().Unix())

		isAt := func(b *big.Int) bool { return b != nil && num.Cmp(b) >= 0 }
		isAtTime := func(t *uint64) bool { return t != nil && ts >= *t }

		n.logger.Println(
			"applyTx: forks",
			"| height:", height,
			"| byz:", isAt(cc.ByzantiumBlock),
			"| cpl:", isAt(cc.ConstantinopleBlock),
			"| ist:", isAt(cc.IstanbulBlock),
			"| ber:", isAt(cc.BerlinBlock),
			"| lon:", isAt(cc.LondonBlock),
			"| sha:", isAtTime(cc.ShanghaiTime),
			"| can:", isAtTime(cc.CancunTime),
			"| ts:", ts,
		)

		// Minimal block/tx context
		const blockGasLimit = uint64(30_000_000)
		baseFee := big.NewInt(1)

		canTransfer := func(db vm.StateDB, addr common.Address, amount *uint256.Int) bool {
			if amount == nil || amount.Sign() == 0 {
				return true
			}
			bal := db.GetBalance(addr)
			return bal.Cmp(amount) >= 0
		}
		doTransfer := func(db vm.StateDB, sender, recipient common.Address, amount *uint256.Int) {
			if amount == nil || amount.Sign() == 0 {
				return
			}
			db.SubBalance(sender, amount, tracing.BalanceChangeUnspecified)
			db.AddBalance(recipient, amount, tracing.BalanceChangeUnspecified)
		}

		blockCtx := vm.BlockContext{
			CanTransfer: canTransfer,
			Transfer:    doTransfer,
			GetHash:     func(uint64) common.Hash { return common.Hash{} },
			Coinbase:    common.Address{},
			GasLimit:    blockGasLimit,
			BlockNumber: new(big.Int).SetUint64(height),
			Time:        uint64(time.Now().Unix()),
			Difficulty:  big.NewInt(0),
			BaseFee:     baseFee,
			Random:      new(common.Hash),
		}
		txCtx := vm.TxContext{Origin: from, GasPrice: tx.GasPrice()}

		evm := vm.NewEVM(blockCtx, txCtx, statedb, &cc, vm.Config{ExtraEips: []int{3855}})

		// Convert tx -> Message and apply transition
		msg, err := core.TransactionToMessage(&tx, signer, baseFee)
		if err != nil {
			n.logger.Println("applyTx: tx->msg failed | hash:", t.Hash, "| err:", err)
			return nil
		}
		gp := new(core.GasPool).AddGas(blockGasLimit)

		res, err := core.ApplyMessage(evm, msg, gp)
		if err != nil {
			n.logger.Println("applyTx: ApplyMessage core-error | hash:", t.Hash, "| err:", err)
		}

		// >>> INSERTION ICI (immédiatement après le if err != nil, avant Finalise)
		if res != nil && res.Err != nil {
			n.logger.Println(
				"applyTx: ApplyMessage FAILED",
				"| hash:", t.Hash,
				"| err:", res.Err,
				"| usedGas:", res.UsedGas,
				"| return:", fmt.Sprintf("0x%x", res.Return),
			)
		}

		// Finalise per-tx to correctly materialize deletions/journal and make logs accessible.
		statedb.Finalise(true)

		// Determine status and usedGas
		status := uint64(1)
		usedGas := uint64(0)
		if res != nil {
			usedGas = res.UsedGas
			if res.Failed() {
				status = 0
			}
		} else {
			status = 0
		}

		// Pull logs from statedb (tx-context aware)
		logs := statedb.GetLogs(rawHash, height, blockHash)

		// Contract creation: derive address for receipt when tx.To()==nil (CREATE)
		// Note: this matches geth semantics (CreateAddress(from, tx.Nonce()))
		var contractAddrPtr *common.Address
		if tx.To() == nil && from != (common.Address{}) {
			ca := crypto.CreateAddress(from, tx.Nonce())
			contractAddrPtr = &ca
		}

		// M9/M13: persist receipt JSON (best-effort) aligned with executed results
		if n.db != nil {
			toPtr := tx.To()
			// Global logIndex within the block (mainnet-like)
			base := uint64(0)
			if blockLogBase != nil {
				base = *blockLogBase
			}

			// Minimal log encoding (enough for tooling) + write-path log index (M14-B)
			logsAny := make([]any, 0, len(logs))
			for i := range logs {
				lg := logs[i]
				topics := make([]string, 0, len(lg.Topics))
				for _, tp := range lg.Topics {
					topics = append(topics, tp.Hex())
				}

				gi := base + uint64(i) // global log index within block
				lm := map[string]any{
					"address":          lg.Address.Hex(),
					"topics":           topics,
					"data":             "0x" + hex.EncodeToString(lg.Data),
					"blockNumber":      toHexUint(height),
					"transactionHash":  rawHash.Hex(),
					"transactionIndex": toHexUint(uint64(txIndex)),
					"blockHash":        blockHash.Hex(),
					"logIndex":         toHexUint(gi),
					"removed":          false,
				}
				logsAny = append(logsAny, lm)

				// logrec/v1/<heightBE><txIndexBE><logIndexBE> -> JSON(logObject)
				if b, err := json.Marshal(lm); err == nil {
					_ = n.db.Put(logRecKey(height, uint64(txIndex), gi), b, nil)
				}
			}
			if blockLogBase != nil {
				*blockLogBase = base + uint64(len(logs))
			}

			rcpt := receiptJSON{
				TransactionHash:   rawHash,
				TransactionIndex:  toHexUint(uint64(txIndex)),
				BlockHash:         blockHash,
				BlockNumber:       toHexUint(height),
				From:              from,
				To:                toPtr,
				CumulativeGasUsed: toHexBig(usedGas),
				GasUsed:           toHexBig(usedGas),
				ContractAddress:   contractAddrPtr,
				Logs:              logsAny,
				Status:            toHexUint(status),
				Type:              toHexUint(uint64(tx.Type())),
			}
			if b, err := json.Marshal(rcpt); err == nil {
				_ = n.db.Put(rcptKey(rawHash.Hex()), b, nil)
			}
		}

		// M12: build a minimal geth receipt for per-block receiptsRoot/logsBloom
		rcptGeth := &types.Receipt{
			TxHash:            rawHash,
			BlockHash:         pseudoBlockHash(height),
			BlockNumber:       new(big.Int).SetUint64(height),
			TransactionIndex:  uint(txIndex),
			Status:            status,
			CumulativeGasUsed: usedGas,
			GasUsed:           usedGas,
			ContractAddress: func() common.Address {
				if contractAddrPtr != nil {
					return *contractAddrPtr
				}
				return common.Address{}
			}(),
			Logs: logs,
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
			"| status:", status,
			"| usedGas:", usedGas,
			"| logs:", len(logs),
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
			blockLogBase := uint64(0)
			if n.txpool != nil && n.txindex != nil {
				txs := n.txpool.PopPending(64)
				for i := range txs {
					n.txindex.Put(txs[i].Hash, height)
					// M9: hook point (execution will be extended in next steps)
					r := applyTx(txs[i], height, i, &blockLogBase)
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
						_ = n.db.Put([]byte(headHeightKVKey), u64be(height), nil)

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
				_ = n.db.Put([]byte(headHeightKVKey), u64be(height), nil)

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

func (n *Node) TxPool() *txpool.Pool      { return n.txpool }
func (n *Node) TxIndex() *txindex.Index   { return n.txindex }
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
