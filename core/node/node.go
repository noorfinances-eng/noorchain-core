package node

import (
	"os"

	"noorchain-evm-l1/core/config"
)

type Logger interface {
	Println(v ...any)
}

type Node struct {
	cfg    config.Config
	logger Logger
}

func New(cfg config.Config) *Node {
	return &Node{
		cfg:    cfg,
		logger: newLogger(),
	}
}

func (n *Node) Start() {
	if err := os.MkdirAll(n.cfg.DataDir, 0o755); err != nil {
		panic(err)
	}
	n.logger.Println("node started")
	n.logger.Println("chain-id:", n.cfg.ChainID)
	n.logger.Println("data-dir:", n.cfg.DataDir)
}

func (n *Node) Stop() {
	n.logger.Println("node stopped")
}
