package node

import (
	"context"
	"os"
	"time"

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
	ctx     context.Context
	cancel  context.CancelFunc
}

func New(cfg config.Config) *Node {
	ctx, cancel := context.WithCancel(context.Background())
	return &Node{
		cfg:     cfg,
		logger:  newLogger(),
		network: network.New(cfg.P2PAddr),
		ctx:     ctx,
		cancel:  cancel,
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

	go n.loop()
}

func (n *Node) loop() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-n.ctx.Done():
			n.logger.Println("node loop stopped")
			return
		case <-ticker.C:
			n.logger.Println("tick")
		}
	}
}

func (n *Node) Stop() {
	n.cancel()
	n.network.Stop()
	n.logger.Println("node stopped")
}
