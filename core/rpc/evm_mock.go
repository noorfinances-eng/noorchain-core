package rpc

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

// EvmMock is a dev-only in-memory tx/receipt store to unblock tooling (Hardhat/Ethers).
// It does NOT execute EVM. It only accepts signed txs, derives tx hash/sender, returns receipts,
// and provides a minimal "state shim" for PoSS read methods via eth_call.
//
// This is strictly for M6 PoSS V0 bring-up.

func computeCreateAddress(from common.Address, nonce uint64) common.Address {
	// Ethereum CREATE address: keccak256(rlp([from, nonce])) last 20 bytes
	enc, _ := rlp.EncodeToBytes([]any{from, nonce})
	h := crypto.Keccak256(enc)
	return common.BytesToAddress(h[12:])
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

// ---- PoSS V0 shim state ----

type possMetaABI struct {
	SnapshotHash [32]byte
	Uri          string
	PeriodStart  uint64
	PeriodEnd    uint64
	PublishedAt  uint64
	Version      uint32
	Publisher    common.Address
}

type possSigABI struct {
	V uint8
	R [32]byte
	S [32]byte
}

type possState struct {
	Snapshots []possMetaABI // 1-indexed for IDs conceptually; stored 0-indexed
}

var possABI abi.ABI
var possSelSubmit [4]byte
var possSelSnapshotCount [4]byte
var possSelLatestSnapshotId [4]byte
var possSelGetSnapshot [4]byte

func init() {
	// Minimal ABI: submitSnapshot(meta,sigs) + view methods used by the client
	const abiJSON = `[
		{"type":"function","name":"submitSnapshot","stateMutability":"nonpayable",
		 "inputs":[
			{"name":"meta","type":"tuple","components":[
				{"name":"snapshotHash","type":"bytes32"},
				{"name":"uri","type":"string"},
				{"name":"periodStart","type":"uint64"},
				{"name":"periodEnd","type":"uint64"},
				{"name":"version","type":"uint32"}
			]},
			{"name":"sigs","type":"tuple[]","components":[
				{"name":"v","type":"uint8"},
				{"name":"r","type":"bytes32"},
				{"name":"s","type":"bytes32"}
			]}
		 ],
		 "outputs":[]
		},
		{"type":"function","name":"snapshotCount","stateMutability":"view",
		 "inputs":[], "outputs":[{"type":"uint256"}]
		},
		{"type":"function","name":"latestSnapshotId","stateMutability":"view",
		 "inputs":[], "outputs":[{"type":"uint256"}]
		},
		{"type":"function","name":"getSnapshot","stateMutability":"view",
		 "inputs":[{"name":"id","type":"uint256"}],
		 "outputs":[{"name":"","type":"tuple","components":[
			{"name":"snapshotHash","type":"bytes32"},
			{"name":"uri","type":"string"},
			{"name":"periodStart","type":"uint64"},
			{"name":"periodEnd","type":"uint64"},
			{"name":"publishedAt","type":"uint64"},
			{"name":"version","type":"uint32"},
			{"name":"publisher","type":"address"}
		 ]}]
		}
	]`

	a, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		panic(err)
	}
	possABI = a

	copy(possSelSubmit[:], a.Methods["submitSnapshot"].ID)
	copy(possSelSnapshotCount[:], a.Methods["snapshotCount"].ID)
	copy(possSelLatestSnapshotId[:], a.Methods["latestSnapshotId"].ID)
	copy(possSelGetSnapshot[:], a.Methods["getSnapshot"].ID)
}

type EvmMock struct {
	mu sync.Mutex

	// nonce per sender
	nonce map[common.Address]uint64

	// tx hash -> receipt
	receipts map[common.Hash]*receiptJSON

	// tx hash -> tx
	txs map[common.Hash]*types.Transaction

	// PoSS shim state per contract address (registry)
	poss map[common.Address]*possState
}

func NewEvmMock() *EvmMock {
	return &EvmMock{
		nonce:    make(map[common.Address]uint64),
		receipts: make(map[common.Hash]*receiptJSON),
		txs:      make(map[common.Hash]*types.Transaction),
		poss:     make(map[common.Address]*possState),
	}
}

func (m *EvmMock) Accounts() []string {
	// node is non-custodial: no unlocked accounts exposed
	return []string{}
}

func (m *EvmMock) GetTransactionCount(addr common.Address) uint64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.nonce[addr]
}

func (m *EvmMock) BumpNonce(addr common.Address, nonce uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	cur := m.nonce[addr]
	if nonce+1 > cur {
		m.nonce[addr] = nonce + 1
	}
}

// SendRawTransaction stores tx + receipt, bumps nonce.
// Additionally, if tx is a call to PoSSRegistry.submitSnapshot(meta,sigs),
// it updates a minimal in-memory PoSS state keyed by recipient contract address.
func (m *EvmMock) SendRawTransaction(rawHex string, chainID *big.Int, blockNumber uint64) (common.Hash, error) {
	rawHex = strings.TrimSpace(rawHex)
	rawHex = strings.TrimPrefix(rawHex, "0x")
	b, err := hex.DecodeString(rawHex)
	if err != nil {
		return common.Hash{}, errors.New("invalid raw tx hex")
	}

	var tx types.Transaction
	if err := tx.UnmarshalBinary(b); err != nil {
		return common.Hash{}, errors.New("invalid raw tx encoding")
	}

	signer := types.LatestSignerForChainID(chainID)
	from, err := types.Sender(signer, &tx)
	if err != nil {
		return common.Hash{}, errors.New("cannot recover sender")
	}

	h := tx.Hash()

	// contract address if creation
	var contractAddr *common.Address
	if tx.To() == nil {
		addr := computeCreateAddress(from, tx.Nonce())
		contractAddr = &addr
	}

	// If this is a PoSS submitSnapshot call, update shim state
	if tx.To() != nil {
		to := *tx.To()
		data := tx.Data()
		if len(data) >= 4 && bytes.Equal(data[:4], possSelSubmit[:]) {
			args := data[4:]
			decoded, derr := possABI.Methods["submitSnapshot"].Inputs.Unpack(args)
			if derr == nil && len(decoded) == 2 {
				// decoded[0] is meta (tuple), decoded[1] is sigs (tuple[])
				// Decode meta tuple explicitly (avoid abi.ConvertType pitfalls)
				metaAny := decoded[0]

				var (
					snapshotHash [32]byte
					uri          string
					periodStart  uint64
					periodEnd    uint64
					version      uint32
				)

				switch v := metaAny.(type) {
				case map[string]any:
					// Most stable path
					if b, ok := v["snapshotHash"].([32]byte); ok {
						snapshotHash = b
					} else if bb, ok := v["snapshotHash"].([]byte); ok && len(bb) == 32 {
						copy(snapshotHash[:], bb)
					}
					uri, _ = v["uri"].(string)

					if ps, ok := v["periodStart"].(uint64); ok {
						periodStart = ps
					} else if ps, ok := v["periodStart"].(*big.Int); ok {
						periodStart = ps.Uint64()
					}
					if pe, ok := v["periodEnd"].(uint64); ok {
						periodEnd = pe
					} else if pe, ok := v["periodEnd"].(*big.Int); ok {
						periodEnd = pe.Uint64()
					}
					if vv, ok := v["version"].(uint32); ok {
						version = vv
					} else if vv, ok := v["version"].(*big.Int); ok {
						version = uint32(vv.Uint64())
					}

				case []any:
					// Positional tuple: [hash, uri, start, end, version]
					if len(v) >= 5 {
						if b, ok := v[0].([32]byte); ok {
							snapshotHash = b
						} else if bb, ok := v[0].([]byte); ok && len(bb) == 32 {
							copy(snapshotHash[:], bb)
						}
						uri, _ = v[1].(string)

						if ps, ok := v[2].(uint64); ok {
							periodStart = ps
						} else if ps, ok := v[2].(*big.Int); ok {
							periodStart = ps.Uint64()
						}
						if pe, ok := v[3].(uint64); ok {
							periodEnd = pe
						} else if pe, ok := v[3].(*big.Int); ok {
							periodEnd = pe.Uint64()
						}
						if vv, ok := v[4].(uint32); ok {
							version = vv
						} else if vv, ok := v[4].(*big.Int); ok {
							version = uint32(vv.Uint64())
						}
					}
				}

				// Fallback: go-ethereum may decode tuples as anonymous structs; extract via reflection
				if version == 0 {
					rv := reflect.ValueOf(metaAny)
					if rv.Kind() == reflect.Struct {
						// SnapshotHash
						fh := rv.FieldByName("SnapshotHash")
						if fh.IsValid() && fh.Kind() == reflect.Array && fh.Len() == 32 {
							for i := 0; i < 32; i++ {
								snapshotHash[i] = byte(fh.Index(i).Uint())
							}
						}
						// Uri
						fu := rv.FieldByName("Uri")
						if fu.IsValid() && fu.Kind() == reflect.String {
							uri = fu.String()
						}
						// PeriodStart / PeriodEnd
						fps := rv.FieldByName("PeriodStart")
						if fps.IsValid() && (fps.Kind() == reflect.Uint64 || fps.Kind() == reflect.Uint32 || fps.Kind() == reflect.Uint) {
							periodStart = uint64(fps.Uint())
						}
						fpe := rv.FieldByName("PeriodEnd")
						if fpe.IsValid() && (fpe.Kind() == reflect.Uint64 || fpe.Kind() == reflect.Uint32 || fpe.Kind() == reflect.Uint) {
							periodEnd = uint64(fpe.Uint())
						}
						// Version
						fv := rv.FieldByName("Version")
						if fv.IsValid() && (fv.Kind() == reflect.Uint32 || fv.Kind() == reflect.Uint64 || fv.Kind() == reflect.Uint) {
							version = uint32(fv.Uint())
						}
					}
				}

				meta := possMetaABI{
					SnapshotHash: snapshotHash,
					Uri:          uri,
					PeriodStart:  periodStart,
					PeriodEnd:    periodEnd,
					PublishedAt:  uint64(time.Now().Unix()),
					Version:      version,
					Publisher:    from,
				}

				// sigs are ignored in shim (signature validated by Solidity in real EVM)

				m.mu.Lock()
				st := m.poss[to]
				if st == nil {
					st = &possState{Snapshots: []possMetaABI{}}
					m.poss[to] = st
				}
				st.Snapshots = append(st.Snapshots, meta)
				m.mu.Unlock()
			}
		}
	}

	// minimal receipt
	rcpt := &receiptJSON{
		TransactionHash:   h,
		TransactionIndex:  "0x0",
		BlockHash:         pseudoBlockHash(blockNumber),
		BlockNumber:       toHexUint(blockNumber),
		From:              from,
		To:                tx.To(),
		CumulativeGasUsed: toHexBig(tx.Gas()),
		GasUsed:           toHexBig(tx.Gas()),
		ContractAddress:   contractAddr,
		Logs:              []any{},
		Status:            "0x1",
		Type:              toHexUint(uint64(tx.Type())),
	}

	m.mu.Lock()
	m.txs[h] = &tx
	m.receipts[h] = rcpt
	m.mu.Unlock()

	m.BumpNonce(from, tx.Nonce())

	// simulate mining delay for polling tools
	time.Sleep(50 * time.Millisecond)

	return h, nil
}

func (m *EvmMock) GetTransactionReceipt(hash common.Hash) *receiptJSON {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.receipts[hash]
}

// Call implements minimal eth_call for PoSS view functions on the PoSSRegistry contract address.
func (m *EvmMock) Call(to common.Address, data []byte) ([]byte, bool) {
	if len(data) < 4 {
		return nil, false
	}
	sel := data[:4]

	m.mu.Lock()
	st := m.poss[to]
	m.mu.Unlock()
	if st == nil {
		// unknown contract => return empty result (client may treat as 0)
		st = &possState{Snapshots: []possMetaABI{}}
	}

	switch {
	case bytes.Equal(sel, possSelSnapshotCount[:]):
		out, _ := possABI.Methods["snapshotCount"].Outputs.Pack(new(big.Int).SetUint64(uint64(len(st.Snapshots))))
		return out, true

	case bytes.Equal(sel, possSelLatestSnapshotId[:]):
		latest := uint64(len(st.Snapshots))
		out, _ := possABI.Methods["latestSnapshotId"].Outputs.Pack(new(big.Int).SetUint64(latest))
		return out, true

	case bytes.Equal(sel, possSelGetSnapshot[:]):
		args := data[4:]
		decoded, err := possABI.Methods["getSnapshot"].Inputs.Unpack(args)
		if err != nil || len(decoded) != 1 {
			return nil, false
		}
		idBig, ok := decoded[0].(*big.Int)
		if !ok {
			return nil, false
		}
		id := idBig.Uint64()

		var meta possMetaABI
		if id >= 1 && id <= uint64(len(st.Snapshots)) {
			meta = st.Snapshots[id-1]
		} else {
			meta = possMetaABI{} // zero
		}

		out, _ := possABI.Methods["getSnapshot"].Outputs.Pack(meta)
		return out, true
	default:
		return nil, false
	}
}

func pseudoBlockHash(n uint64) common.Hash {
	// deterministic pseudo hash for dev
	b := make([]byte, 32)
	for i := 0; i < 8; i++ {
		b[31-i] = byte(n >> (8 * i))
	}
	return common.BytesToHash(crypto.Keccak256(b))
}

func toHexBig(v uint64) string {
	return "0x" + new(big.Int).SetUint64(v).Text(16)
}

// JSON helpers
func mustJSON(v any) json.RawMessage {
	b, _ := json.Marshal(v)
	return b
}
