package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdkcodec "github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkstd "github.com/cosmos/cosmos-sdk/std"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
)

// EncodingConfig contient les objets d'encodage utilisés par NOORCHAIN.
//
// C'est le schéma standard Cosmos SDK :
// - InterfaceRegistry : enregistre les interfaces et implémentations
// - Marshaler (Codec) : codec Protobuf pour (dé)sérialiser les messages
// - TxConfig : configuration pour construire / encoder / décoder les transactions
// - Amino : ancien codec (encore utilisé dans certains cas)
type EncodingConfig struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Marshaler         sdkcodec.Codec
	TxConfig          client.TxConfig
	Amino             *sdkcodec.LegacyAmino
}

// MakeEncodingConfig crée une configuration d'encodage réelle pour NOORCHAIN.
//
// Pour l'instant :
// - on crée les codecs standards Cosmos
// - on enregistre les types "std" (banque, gov, etc.)
// - on n'enregistre pas encore les modules NOOR spécifiques (PoSS, EVM...)
func MakeEncodingConfig() EncodingConfig {
	// 1) Ancien codec Amino (encore utile pour certains cas)
	amino := sdkcodec.NewLegacyAmino()

	// 2) Registry des interfaces (Msg, Account, etc.)
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	// 3) Codec Protobuf principal
	marshaler := sdkcodec.NewProtoCodec(interfaceRegistry)

	// 4) Configuration des transactions (builder, encoder, decoder)
	txCfg := authtx.NewTxConfig(marshaler, authtx.DefaultSignModes)

	// 5) Construire la config
	encCfg := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}

	// 6) Enregistrer les types standards Cosmos (banque, staking, gov, etc.)
	//    Plus tard, on ajoutera ici l'enregistrement des modules NOOR (PoSS, EVM…)
	sdkstd.RegisterLegacyAminoCodec(encCfg.Amino)
	sdkstd.RegisterInterfaces(encCfg.InterfaceRegistry)

	return encCfg
}
