package node

import (
	"fmt"

	"noorchain-evm-l1/core/config"
)

type Node struct {
	cfg config.Config
}

func New(cfg config.Config) *Node {
	return &Node{cfg: cfg}
}

func (n *Node) Start() {
	fmt.Println("node started")
	fmt.Println("chain-id:", n.cfg.ChainID)
	fmt.Println("data-dir:", n.cfg.DataDir)
}

func (n *Node) Stop() {
	fmt.Println("node stopped")
}
