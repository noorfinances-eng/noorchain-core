package app

// NOTE:
// This file will later contain the encoding configuration for NOORCHAIN,
// wiring together the Protobuf and JSON codecs required by the Cosmos SDK
// and Ethermint (for EVM compatibility).
//
// For now, we only define minimal placeholder types and functions so that
// the project structure is clear and ready to be extended.

type EncodingConfig struct {
    // In a full Cosmos SDK app, this struct usually contains the following:
    // - InterfaceRegistry
    // - Codec (Proto codec)
    // - TxConfig
    // - Amino (legacy codec)
    //
    // We keep it empty for now and will replace it when we wire Cosmos SDK.
}

// MakeEncodingConfig returns a placeholder encoding configuration.
// In the next phases, it will be updated to build a real Cosmos SDK
// encoding configuration, including all NOORCHAIN modules.
func MakeEncodingConfig() EncodingConfig {
    return EncodingConfig{}
}
