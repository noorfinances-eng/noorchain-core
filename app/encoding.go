package app

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

// EncodingConfig regroupe tous les encodeurs / decodeurs nécessaires
// pour les modules Cosmos SDK.
type EncodingConfig struct {
	Amino             *codec.LegacyAmino
	InterfaceRegistry codectypes.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          codec.TxConfig
}

// MakeEncodingConfig initialise tous les encodeurs Cosmos SDK.
//
// NOTE : pour l'instant, aucun module personnalisé n'est ajouté ici.
// Lors de la Phase 4, le module PoSS sera enregistré.
func MakeEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := codec.NewTxConfig(marshaler, codec.DefaultSigningModes())

	return EncodingConfig{
		Amino:             amino,
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
	}
}
