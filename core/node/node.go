package node

import (
	"os"

	"noorchain-evm-l1/core/config"
	"noorchain-evm-l1/core/network"
)

type Logger interface {
	Println(v ...any)
}

type Node struct {
	cfg     config.Config
	logger  Logger
	network *network.Network
}

func New(cfg config.Config) *Node {
	return &Node{
		cfg:     cfg,
		logger:  newLogger(),
		network: network.New("127.0.0.1:30303"),
	}
}

func (n *Node) Start() {
	if err := os.MkdirAll(n.cfg.DataDir, 0o755); err != nil {
		panic(err)
	}
	n.logger.Println("node started")
	n.logger.Println("chain-id:", n.cfg.ChainID)
	n.logger.Println("data-dir:", n.cfg.DataDir)

	n.network.Start()
}

func (n *Node) Stop() {
	n.network.Stop()
	n.logger.Println("node stopped")
}
