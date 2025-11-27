package app

import (
    "github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

    "github.com/cosmos/cosmos-sdk/types/module"
)

// EncodingConfig holds all encoding-related types for NOORCHAIN.
type EncodingConfig struct {
    InterfaceRegistry cdctypes.InterfaceRegistry
    Codec             codec.Codec
    Amino             *codec.LegacyAmino
    TxConfig          codec.TxConfig
}

// MakeEncodingConfig builds the encoding configuration for the app.
func MakeEncodingConfig() EncodingConfig {
    amino := codec.NewLegacyAmino()
    interfaceRegistry := cdctypes.NewInterfaceRegistry()
    codec := codec.NewProtoCodec(interfaceRegistry)

    txConfig := codec.NewTxConfig(codec, amino)

    // Register all modules’ interfaces
    ModuleBasics.RegisterLegacyAminoCodec(amino)
    ModuleBasics.RegisterInterfaces(interfaceRegistry)

    return EncodingConfig{
        InterfaceRegistry: interfaceRegistry,
        Codec:             codec,
        Amino:             amino,
        TxConfig:          txConfig,
    }
}
