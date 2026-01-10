package config

type Config struct {
        AllocFile  string
        ChainID     string
        DataDir     string
        P2PAddr     string
        BootPeers   []string
        Role        string
        FollowRPC   string
        HealthAddr  string
}

func Default() Config {
        return Config{
                ChainID:     "noorchain-2-1-local",
                DataDir:     "./data",
                P2PAddr:     "127.0.0.1:30303",
                Role:        "leader",
                FollowRPC:   "",
                AllocFile:  "",
                HealthAddr:  "127.0.0.1:8080",
        }
}
