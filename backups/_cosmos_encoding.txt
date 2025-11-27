package app

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// CosmosEncodingConfig defines the real Cosmos encoding configuration.
// In Phase 2, this is only a placeholder structure matching Cosmos expectations.
type CosmosEncodingConfig struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          codec.TxConfig
	Amino             *codec.LegacyAmino
}

// MakeCosmosEncodingConfig creates an empty CosmosEncodingConfig placeholder.
func MakeCosmosEncodingConfig() CosmosEncodingConfig {
	ir := codectypes.NewInterfaceRegistry()

	return CosmosEncodingConfig{
		InterfaceRegistry: ir,
		Marshaler:         codec.NewProtoCodec(ir),
		TxConfig:          nil,
		Amino:             codec.NewLegacyAmino(),
	}
}
