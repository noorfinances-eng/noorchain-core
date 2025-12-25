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

type State string

const (
	StateInit     State = "INIT"
	StateRunning  State = "RUNNING"
	StateStopping State = "STOPPING"
)

type Node struct {
	cfg     config.Config
	logger  Logger
	network *network.Network
	ctx     context.Context
	cancel  context.CancelFunc
	state   State
	height  uint64
}

func New(cfg config.Config) *Node {
	ctx, cancel := context.WithCancel(context.Background())
	return &Node{
		cfg:     cfg,
		logger:  newLogger(),
		network: network.New(cfg.P2PAddr),
		ctx:     ctx,
		cancel:  cancel,
		state:   StateInit,
		height:  0,
	}
}

func (n *Node) Start() {
	if err := os.MkdirAll(n.cfg.DataDir, 0o755); err != nil {
		panic(err)
	}

	n.state = StateRunning
	n.logger.Println("node started")
	n.logger.Println("state:", n.state)
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
			n.height++
			n.logger.Println("tick | height:", n.height, "| state:", n.state)
		}
	}
}

func (n *Node) Stop() {
	n.state = StateStopping
	n.logger.Println("state:", n.state)

	n.cancel()
	n.network.Stop()
	n.logger.Println("node stopped")
}
