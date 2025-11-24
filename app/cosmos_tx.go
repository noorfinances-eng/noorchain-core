package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CosmosTxDecoder is a placeholder transaction decoder for NOORCHAIN.
// In Phase 2 this returns nil and will be replaced by a real decoder later.
func CosmosTxDecoder(txBytes []byte) (sdk.Tx, error) {
	_ = txBytes
	return nil, nil
}
