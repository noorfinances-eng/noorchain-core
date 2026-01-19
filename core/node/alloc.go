package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

const allocAppliedKVKey = "alloc/v1/applied"

type allocFile struct {
	ChainID uint64       `json:"chainId"`
	Alloc   []allocEntry `json:"alloc"`
}

type allocEntry struct {
	Address    string `json:"address"`
	BalanceWei string `json:"balanceWei"` // decimal string
}

func readAllocFile(path string) (*allocFile, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var af allocFile
	if err := json.Unmarshal(b, &af); err != nil {
		return nil, err
	}
	if af.ChainID == 0 {
		return nil, errors.New("alloc: chainId missing/zero")
	}
	if len(af.Alloc) == 0 {
		return nil, errors.New("alloc: empty alloc list")
	}
	return &af, nil
}

func parseAllocEntries(af *allocFile) (map[common.Address]*uint256.Int, error) {
	out := make(map[common.Address]*uint256.Int, len(af.Alloc))
	for i, e := range af.Alloc {
		addrStr := strings.TrimSpace(e.Address)
		if !common.IsHexAddress(addrStr) {
			return nil, fmt.Errorf("alloc: invalid address at index %d", i)
		}
		a := common.HexToAddress(addrStr)
		if _, exists := out[a]; exists {
			return nil, fmt.Errorf("alloc: duplicate address %s", a.Hex())
		}

		balStr := strings.TrimSpace(e.BalanceWei)
		if balStr == "" {
			return nil, fmt.Errorf("alloc: empty balanceWei at index %d", i)
		}
		v, err := uint256.FromDecimal(balStr)
		if err != nil {
			return nil, fmt.Errorf("alloc: invalid balanceWei at index %d", i)
		}
		out[a] = v
	}
	return out, nil
}
