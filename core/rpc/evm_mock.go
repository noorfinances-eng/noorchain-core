package rpc

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

// EvmMock is a dev-only tx/receipt store to unblock tooling (Hardhat/Ethers).
// It does NOT execute EVM. It accepts signed txs, derives tx hash/sender, returns receipts,
// and provides a minimal "state shim" for PoSS read methods via eth_call.
//
// PoSS shim state is persisted in LevelDB so it survives restarts.
//
// ---- KV layout ----
// poss/v1/<registryAddrLowerHex>/count              -> uint64 (big-endian)
// poss/v1/<registryAddrLowerHex>/latest             -> uint64 (big-endian)
// poss/v1/<registryAddrLowerHex>/snap/<id_u64_be>   -> rlp(possMetaABI)

const possKVPrefix = "poss/v1/"

func possAddrPrefix(addr common.Address) []byte {
	return []byte(possKVPrefix + strings.ToLower(addr.Hex()) + "/")
}

func possKeyCount(addr common.Address) []byte {
	return append(possAddrPrefix(addr), []byte("count")...)
}
func possKeyLatest(addr common.Address) []byte {
	return append(possAddrPrefix(addr), []byte("latest")...)
}

func possKeySnap(addr common.Address, id uint64) []byte {
	p := possAddrPrefix(addr)
	p = append(p, []byte("snap/")...)
	var be [8]byte
	binary.BigEndian.PutUint64(be[:], id)
	p = append(p, be[:]...)
	return p
}

func u64ToBE(v uint64) []byte {
	var be [8]byte
	binary.BigEndian.PutUint64(be[:], v)
	return be[:]
}

func beToU64(b []byte) (uint64, bool) {
	if len(b) != 8 {
		return 0, false
	}
	return binary.BigEndian.Uint64(b), true
}

func computeCreateAddress(from common.Address, nonce uint64) common.Address {
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
	Snapshots []possMetaABI
}

var possABI abi.ABI
var possSelSubmit [4]byte
var possSelSnapshotCount [4]byte
var possSelLatestSnapshotId [4]byte
var possSelGetSnapshot [4]byte

func init() {
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
	db *leveldb.DB
	mu sync.Mutex

	nonce    map[common.Address]uint64
	receipts map[common.Hash]*receiptJSON
	txs      map[common.Hash]*types.Transaction

	// optional cache; DB is source of truth for reads
	poss map[common.Address]*possState
}

func NewEvmMock(db *leveldb.DB) *EvmMock {
	return &EvmMock{
		db:       db,
		nonce:    make(map[common.Address]uint64),
		receipts: make(map[common.Hash]*receiptJSON),
		txs:      make(map[common.Hash]*types.Transaction),
		poss:     make(map[common.Address]*possState),
	}
}

func (m *EvmMock) Accounts() []string { return []string{} }

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

// ---- PoSS persistence helpers ----

func (m *EvmMock) possPersistSnapshot(to common.Address, meta possMetaABI) {
	if m.db == nil {
		return
	}

	var curCount uint64
	if b, err := m.db.Get(possKeyCount(to), nil); err == nil {
		if v, ok := beToU64(b); ok {
			curCount = v
		}
	}
	newID := curCount + 1

	enc, err := rlp.EncodeToBytes(meta)
	if err == nil {
		_ = m.db.Put(possKeySnap(to, newID), enc, nil)
	}
	_ = m.db.Put(possKeyCount(to), u64ToBE(newID), nil)
	_ = m.db.Put(possKeyLatest(to), u64ToBE(newID), nil)
}

func (m *EvmMock) possReadCount(to common.Address) uint64 {
	if m.db == nil {
		return 0
	}
	b, err := m.db.Get(possKeyCount(to), nil)
	if err != nil {
		return 0
	}
	v, ok := beToU64(b)
	if !ok {
		return 0
	}
	return v
}

func (m *EvmMock) possReadLatest(to common.Address) uint64 {
	if m.db == nil {
		return 0
	}
	b, err := m.db.Get(possKeyLatest(to), nil)
	if err != nil {
		return 0
	}
	v, ok := beToU64(b)
	if !ok {
		return 0
	}
	return v
}

func (m *EvmMock) possReadSnapshot(to common.Address, id uint64) (possMetaABI, bool) {
	if m.db == nil || id == 0 {
		return possMetaABI{}, false
	}
	b, err := m.db.Get(possKeySnap(to, id), nil)
	if err != nil || len(b) == 0 {
		return possMetaABI{}, false
	}
	var meta possMetaABI
	if err := rlp.DecodeBytes(b, &meta); err != nil {
		return possMetaABI{}, false
	}
	return meta, true
}

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

	var contractAddr *common.Address
	if tx.To() == nil {
		addr := computeCreateAddress(from, tx.Nonce())
		contractAddr = &addr
	}

	if tx.To() != nil {
		to := *tx.To()
		data := tx.Data()
		if len(data) >= 4 {
			args := data[4:]
			decoded, derr := possABI.Methods["submitSnapshot"].Inputs.Unpack(args)
			if derr == nil && len(decoded) == 2 {
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
					if bb, ok := v["snapshotHash"].([32]byte); ok {
						snapshotHash = bb
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
					if len(v) >= 5 {
						if bb, ok := v[0].([32]byte); ok {
							snapshotHash = bb
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

				if version == 0 {
					rv := reflect.ValueOf(metaAny)
					if rv.Kind() == reflect.Struct {
						fh := rv.FieldByName("SnapshotHash")
						if fh.IsValid() && fh.Kind() == reflect.Array && fh.Len() == 32 {
							for i := 0; i < 32; i++ {
								snapshotHash[i] = byte(fh.Index(i).Uint())
							}
						}
						fu := rv.FieldByName("Uri")
						if fu.IsValid() && fu.Kind() == reflect.String {
							uri = fu.String()
						}
						fps := rv.FieldByName("PeriodStart")
						if fps.IsValid() {
							periodStart = uint64(fps.Uint())
						}
						fpe := rv.FieldByName("PeriodEnd")
						if fpe.IsValid() {
							periodEnd = uint64(fpe.Uint())
						}
						fv := rv.FieldByName("Version")
						if fv.IsValid() {
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

				m.mu.Lock()
				st := m.poss[to]
				if st == nil {
					st = &possState{Snapshots: []possMetaABI{}}
					m.poss[to] = st
				}
				st.Snapshots = append(st.Snapshots, meta)
				m.mu.Unlock()

				m.possPersistSnapshot(to, meta)
			}
		}
	}

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
	time.Sleep(50 * time.Millisecond)

	return h, nil
}

func (m *EvmMock) GetTransactionReceipt(hash common.Hash) *receiptJSON {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.receipts[hash]
}

func (m *EvmMock) Call(to common.Address, data []byte) ([]byte, bool) {
	if len(data) < 4 {
		return nil, false
	}
	sel := data[:4]

	if m.db != nil {
		switch {
		case bytes.Equal(sel, possSelSnapshotCount[:]):
			c := m.possReadCount(to)
			out, _ := possABI.Methods["snapshotCount"].Outputs.Pack(new(big.Int).SetUint64(c))
			return out, true

		case bytes.Equal(sel, possSelLatestSnapshotId[:]):
			latest := m.possReadLatest(to)
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

			meta, ok := m.possReadSnapshot(to, id)
			if !ok {
				meta = possMetaABI{}
			}
			out, _ := possABI.Methods["getSnapshot"].Outputs.Pack(meta)
			return out, true
		}
	}

	m.mu.Lock()
	st := m.poss[to]
	m.mu.Unlock()
	if st == nil {
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
			meta = possMetaABI{}
		}
		out, _ := possABI.Methods["getSnapshot"].Outputs.Pack(meta)
		return out, true

	default:
		return nil, false
	}
}

func pseudoBlockHash(n uint64) common.Hash {
	b := make([]byte, 32)
	for i := 0; i < 8; i++ {
		b[31-i] = byte(n >> (8 * i))
	}
	return common.BytesToHash(crypto.Keccak256(b))
}

func toHexBig(v uint64) string {
	return "0x" + new(big.Int).SetUint64(v).Text(16)
}

func mustJSON(v any) json.RawMessage {
	b, _ := json.Marshal(v)
	return b
}
