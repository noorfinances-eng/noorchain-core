package app

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// EncodingConfig rassemble tous les encodeurs / décodeurs nécessaires
// au fonctionnement de NOORCHAIN. Ce fichier est stable et compatible
// Cosmos SDK v0.46.x.
type EncodingConfig struct {
	Amino             *codec.LegacyAmino
	InterfaceRegistry codectypes.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          codec.TxConfig
}

// MakeEncodingConfig initialise tous les encodeurs Cosmos SDK.
// Aucun module personnalisé n'est enregistré ici pour l’instant.
// (Le module PoSS sera ajouté en Phase 4.)
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
