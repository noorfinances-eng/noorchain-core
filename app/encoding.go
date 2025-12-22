package app

import (
	sigopts "cosmossdk.io/x/tx/signing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/std"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
)

// EncodingConfig garde le codec et la config Tx utilisés par l'app/CLI.
type EncodingConfig struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// MakeEncodingConfig construit une config standard Cosmos SDK v0.53
// avec :
// - SigningOptions (AddressCodec/ValidatorAddressCodec) adaptés à NOOR,
// - Types des modules de l'app (ModuleBasics) enregistrés,
// - TxConfig classique (auth/tx).
func MakeEncodingConfig() EncodingConfig {
	// Bech32 codecs pour NOORCHAIN
	accAddrCodec := addresscodec.NewBech32Codec("noor")
	valAddrCodec := addresscodec.NewBech32Codec("noorvaloper")

	// Options de signature : indispensables pour éviter
	// "InterfaceRegistry requires a proper address codec implementation..."
	signingOptions := sigopts.Options{
		AddressCodec:          accAddrCodec,
		ValidatorAddressCodec: valAddrCodec,
		// ConsensusAddressCodec peut rester nil à ce stade M2.
	}

	// InterfaceRegistry avec SigningOptions
	interfaceRegistry, err := codectypes.NewInterfaceRegistryWithOptions(
		codectypes.InterfaceRegistryOptions{
			SigningOptions: signingOptions,
		},
	)
	if err != nil {
		panic(err)
	}

	// 1) Types SDK de base
	std.RegisterInterfaces(interfaceRegistry)

	// 2) Types crypto (PubKey, etc.)
	cryptocodec.RegisterInterfaces(interfaceRegistry)

	// 3) Types des modules de l'app (auth, bank, genutil)
	ModuleBasics.RegisterInterfaces(interfaceRegistry)

	// Codec Proto
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	// TxConfig standard (auth/tx)
	txCfg := authtx.NewTxConfig(marshaler, authtx.DefaultSignModes)

	// Legacy Amino + modules (utile pour certains outils CLI)
	amino := codec.NewLegacyAmino()
	std.RegisterLegacyAminoCodec(amino)
	ModuleBasics.RegisterLegacyAminoCodec(amino)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}
