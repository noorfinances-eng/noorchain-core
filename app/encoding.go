package app

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

// EncodingConfig defines the application encoding configuration.
// It is used to initialize all codec-related features.
type EncodingConfig struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// MakeEncodingConfig creates an EncodingConfig for testing or app setup.
func MakeEncodingConfig(modules module.BasicManager) EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txConfig := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	// Register modules interfaces & codecs
	modules.RegisterLegacyAminoCodec(amino)
	modules.RegisterInterfaces(interfaceRegistry)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txConfig,
		Amino:             amino,
	}
}
