package node

import (
	"fmt"
	"os"

	"noorchain-evm-l1/core/config"
)

type Node struct {
	cfg config.Config
}

func New(cfg config.Config) *Node {
	return &Node{cfg: cfg}
}

func (n *Node) Start() {
	if err := os.MkdirAll(n.cfg.DataDir, 0o755); err != nil {
		panic(err)
	}
	fmt.Println("node started")
	fmt.Println("chain-id:", n.cfg.ChainID)
	fmt.Println("data-dir:", n.cfg.DataDir)
}

func (n *Node) Stop() {
	fmt.Println("node stopped")
}
