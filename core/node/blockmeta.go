package node

import (
	"encoding/json"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// ---- M12 block metadata persistence (stateRoot/receiptsRoot/logsBloom) ----
//
// Key layout (in NOOR LevelDB <data-dir>/db/leveldb):
//   blkmeta/v1/<height_u64_hex_no0x>  -> json(blockMeta)

const blkMetaPrefix = "blkmeta/v1/"

type blockMeta struct {
	Height       uint64      `json:"height"`
	Timestamp    uint64      `json:"timestamp"` // unix seconds
	BlockHash    common.Hash `json:"blockHash"`
	StateRoot    common.Hash `json:"stateRoot"`
	TransactionsRoot common.Hash `json:"transactionsRoot"`
	ReceiptsRoot common.Hash `json:"receiptsRoot"`
	LogsBloomHex string      `json:"logsBloom"` // 0x + 512 hex chars (256 bytes)
}

func blkMetaKey(height uint64) []byte {
	// hex without 0x, lowercase, no padding needed
	h := strings.TrimPrefix(toHexUint(height), "0x")
	return []byte(blkMetaPrefix + h)
}

func encodeBlockMeta(m blockMeta) ([]byte, error) { return json.Marshal(m) }
func decodeBlockMeta(b []byte) (blockMeta, error) {
	var m blockMeta
	err := json.Unmarshal(b, &m)
	return m, err
}
