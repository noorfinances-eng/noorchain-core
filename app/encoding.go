package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdkcodec "github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// EncodingConfig holds the concrete encoding types used by NOORCHAIN.
//
// This is the standard pattern used by modern Cosmos SDK chains:
// - InterfaceRegistry: used to register interfaces and implementations
// - Marshaler (Codec): protobuf codec used to (de)serialize messages
// - TxConfig: configuration and encoder/decoder for transactions
// - Amino: legacy codec (optional but still often kept for backwards compatibility)
type EncodingConfig struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Marshaler         sdkcodec.Codec
	TxConfig          client.TxConfig
	Amino             *sdkcodec.LegacyAmino
}

// MakeEncodingConfig returns a placeholder encoding configuration.
//
// IMPORTANT:
// - For now, all fields are nil / zero values.
// - In next phases, this function will be updated to:
//   - build a proper InterfaceRegistry
//   - create a Protobuf codec
//   - create a TxConfig
//   - register all NOORCHAIN modules (including PoSS)
func MakeEncodingConfig() EncodingConfig {
	return EncodingConfig{
		InterfaceRegistry: nil,
		Marshaler:         nil,
		TxConfig:          nil,
		Amino:             nil,
	}
}
