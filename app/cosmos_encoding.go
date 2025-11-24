package app

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
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
	// Placeholder: in the next steps we will initialize the real registries/codecs.
	return CosmosEncodingConfig{
		InterfaceRegistry: codectypes.NewInterfaceRegistry(),
		Marshaler:         codec.NewProtoCodec(codectypes.NewInterfaceRegistry()),
		TxConfig:          nil,
		Amino:             codec.NewLegacyAmino(),
	}
}
