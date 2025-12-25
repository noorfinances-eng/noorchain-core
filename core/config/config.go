package config

type Config struct {
	ChainID string
	DataDir string
	P2PAddr string
}

func Default() Config {
	return Config{
		ChainID: "noorchain-2-1-local",
		DataDir: "./data",
		P2PAddr: "127.0.0.1:30303",
	}
}
