package app

import (
	"github.com/cosmos/gogoproto/proto"

	txsigning "cosmossdk.io/x/tx/signing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
)

// EncodingConfig bundles the app-wide encoding settings.
type EncodingConfig struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// MakeEncodingConfig builds an EncodingConfig for Noorchain (M2/M3 minimal)
// with Cosmos SDK v0.53.5 and base modules declared in ModuleBasics
// (auth, bank, staking, genutil).
func MakeEncodingConfig() EncodingConfig {
	ir, err := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: txsigning.Options{
			FileResolver:          proto.HybridResolver,
			AddressCodec:          addresscodec.NewBech32Codec("noor"),
			ValidatorAddressCodec: addresscodec.NewBech32Codec("noorvaloper"),
		},
	})
	if err != nil {
		panic(err)
	}

	amino := codec.NewLegacyAmino()
	cdc := codec.NewProtoCodec(ir)

	txCfg, err := authtx.NewTxConfigWithOptions(cdc, authtx.ConfigOptions{
		EnabledSignModes: authtx.DefaultSignModes,
		SigningOptions: &txsigning.Options{
			FileResolver:          proto.HybridResolver,
			AddressCodec:          addresscodec.NewBech32Codec("noor"),
			ValidatorAddressCodec: addresscodec.NewBech32Codec("noorvaloper"),
		},
	})
	if err != nil {
		panic(err)
	}

	std.RegisterInterfaces(ir)
	ModuleBasics.RegisterInterfaces(ir)

	std.RegisterLegacyAminoCodec(amino)
	ModuleBasics.RegisterLegacyAminoCodec(amino)

	return EncodingConfig{
		InterfaceRegistry: ir,
		Marshaler:         cdc,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}
