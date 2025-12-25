package config

type Config struct {
	ChainID string
	DataDir string
}

func Default() Config {
	return Config{
		ChainID: "noorchain-2-1-local",
		DataDir: "./data",
	}
}
