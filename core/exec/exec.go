package exec

import (
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb"
	"math/big"
	"reflect"
	"strconv"
)

// local copy (must match rpc.possMetaABI field order for RLP compatibility)
type possMetaABI struct {
	SnapshotHash [32]byte
	Uri          string
	PeriodStart  uint64
	PeriodEnd    uint64
	PublishedAt  uint64
	Version      uint32
	Publisher    common.Address
}

// ApplyPoSSSubmitSnapshot persists PoSS submitSnapshot calls into NOOR LevelDB.
//
// KV layout (aligned with rpc/evm_mock.go):
//
//	poss/v1/<registryAddrLowerHex>/count              -> uint64 (big-endian)
//	poss/v1/<registryAddrLowerHex>/latest             -> uint64 (big-endian)
//	poss/v1/<registryAddrLowerHex>/snap/<id_u64_be>   -> rlp(possMetaABI)
//
// IMPORTANT: submitSnapshot has NO explicit snapshotId in calldata.
// We assign id deterministically as (count+1) and bump count/latest accordingly.
//
// IMPORTANT: registry address is the tx "to" (no hardcoded dev address).
func ApplyPoSSSubmitSnapshot(tx *types.Transaction, chainID string, db *leveldb.DB, publishedAt uint64) (bool, error) {
	if tx == nil || db == nil {
		return false, nil
	}
	toPtr := tx.To()
	if toPtr == nil {
		return false, nil
	}
	registry := *toPtr

	data := tx.Data()
	if len(data) < 4 {
		return false, nil
	}

	// submitSnapshot selector (sync with rpc/evm_mock.go possSelSubmit)
	submitSel := [4]byte{0x2b, 0xa1, 0x8a, 0x99}
	var gotSel [4]byte
	copy(gotSel[:], data[:4])
	if gotSel != submitSel {
		return false, nil
	}

	regLower := strings.ToLower(registry.Hex())
	base := []byte("poss/v1/" + regLower + "/")

	kCount := append(append([]byte{}, base...), []byte("count")...)
	kLatest := append(append([]byte{}, base...), []byte("latest")...)

	// Load current count/latest.
	var count uint64
	if b, err := db.Get(kCount, nil); err == nil && len(b) == 8 {
		count = binary.BigEndian.Uint64(b)
	} else if err != nil && err != leveldb.ErrNotFound {
		return true, fmt.Errorf("poss: read count: %w", err)
	}

	var latest uint64
	if b, err := db.Get(kLatest, nil); err == nil && len(b) == 8 {
		latest = binary.BigEndian.Uint64(b)
	} else if err != nil && err != leveldb.ErrNotFound {
		return true, fmt.Errorf("poss: read latest: %w", err)
	}

	// Assign new id as count+1 (deterministic).
	newID := count + 1

	// snap key
	var idBE [8]byte
	binary.BigEndian.PutUint64(idBE[:], newID)
	kSnap := append(append(append([]byte{}, base...), []byte("snap/")...), idBE[:]...)

	// If already exists, treat as handled (idempotent) and ensure latest>=newID.
	if _, err := db.Get(kSnap, nil); err == nil {
		if latest < newID {
			latest = newID
			var outLatest [8]byte
			binary.BigEndian.PutUint64(outLatest[:], latest)
			if err := db.Put(kLatest, outLatest[:], nil); err != nil {
				return true, fmt.Errorf("poss: update latest: %w", err)
			}
		}
		return true, nil
	} else if err != nil && err != leveldb.ErrNotFound {
		return true, fmt.Errorf("poss: read snap: %w", err)
	}

	// Decode calldata meta (tuple) so getSnapshot returns full fields.
	//
	// submitSnapshot(meta=(snapshotHash,uri,periodStart,periodEnd,version), sigs=[...])
	const possSubmitABIJSON = `[
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
  }
]`

	var (
		snapshotHash [32]byte
		uri          string
		periodStart  uint64
		periodEnd    uint64
		version      uint32
	)

	// Recover publisher (tx sender).
	publisher := common.Address{}
	evmCID := tx.ChainId()
	if evmCID == nil {
		// best-effort fallback if caller passed numeric chainID
		if n, err := strconv.ParseInt(chainID, 10, 64); err == nil {
			evmCID = big.NewInt(n)
		} else {
			evmCID = big.NewInt(0)
		}
	}
	signer := types.LatestSignerForChainID(evmCID)
	if from, err := types.Sender(signer, tx); err == nil {
		publisher = from
	}

	// ABI decode of meta tuple from calldata (best-effort; fallback keeps compatibility).
	if a, err := abi.JSON(strings.NewReader(possSubmitABIJSON)); err == nil {
		if decoded, derr := a.Methods["submitSnapshot"].Inputs.Unpack(data[4:]); derr == nil && len(decoded) == 2 {
			metaAny := decoded[0]

			switch v := metaAny.(type) {
			case map[string]any:
				if bb, ok := v["snapshotHash"].([32]byte); ok {
					snapshotHash = bb
				} else if bb, ok := v["snapshotHash"].([]byte); ok && len(bb) == 32 {
					copy(snapshotHash[:], bb)
				}
				if u, ok := v["uri"].(string); ok {
					uri = u
				}
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
					if u, ok := v[1].(string); ok {
						uri = u
					}
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

			default:
				// Reflect fallback (tuple decoded as anonymous struct)
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
		}
	}

	// Fallback if snapshotHash still empty: use tx hash (keeps old behavior as safety net).
	var zero32 [32]byte
	if snapshotHash == zero32 {
		h := tx.Hash().Bytes()
		copy(snapshotHash[:], h)
	}

	meta := possMetaABI{
		SnapshotHash: snapshotHash,
		Uri:          uri,
		PeriodStart:  periodStart,
		PeriodEnd:    periodEnd,
		PublishedAt:  publishedAt,
		Version:      version,
		Publisher:    publisher,
	}

	metaRLP, err := rlp.EncodeToBytes(meta)
	if err != nil {
		return true, fmt.Errorf("poss: rlp encode: %w", err)
	}

	// Bump count/latest.
	newCount := newID
	newLatest := latest
	if newLatest < newID {
		newLatest = newID
	}

	var outCount [8]byte
	var outLatest [8]byte
	binary.BigEndian.PutUint64(outCount[:], newCount)
	binary.BigEndian.PutUint64(outLatest[:], newLatest)

	batch := new(leveldb.Batch)
	batch.Put(kCount, outCount[:])
	batch.Put(kLatest, outLatest[:])
	batch.Put(kSnap, metaRLP)

	if err := db.Write(batch, nil); err != nil {
		return true, fmt.Errorf("poss: db write: %w", err)
	}

	return true, nil
}
