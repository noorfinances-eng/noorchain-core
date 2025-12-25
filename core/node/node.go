package node

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"noorchain-evm-l1/core/config"
	"noorchain-evm-l1/core/health"
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
	health  *health.Server

	ctx    context.Context
	cancel context.CancelFunc

	mu     sync.Mutex
	state  State
	height uint64
}

func New(cfg config.Config) *Node {
	ctx, cancel := context.WithCancel(context.Background())
	return &Node{
		cfg:     cfg,
		logger:  newLogger(),
		network: network.New(cfg.P2PAddr),
		health:  health.New("127.0.0.1:8080"),
		ctx:     ctx,
		cancel:  cancel,
		state:   StateInit,
		height:  0,
	}
}

func (n *Node) Start() error {
	if err := os.MkdirAll(n.cfg.DataDir, 0o755); err != nil {
		return fmt.Errorf("mkdir data-dir %s: %w", n.cfg.DataDir, err)
	}

	n.mu.Lock()
	n.state = StateRunning
	n.mu.Unlock()

	n.logger.Println("node started")
	n.logger.Println("state:", StateRunning)
	n.logger.Println("chain-id:", n.cfg.ChainID)
	n.logger.Println("data-dir:", n.cfg.DataDir)

	if err := n.network.Start(); err != nil {
		return err
	}
	if err := n.health.Start(); err != nil {
		return err
	}

	go n.loop()
	return nil
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
			n.mu.Lock()
			n.height++
			state := n.state
			height := n.height
			n.mu.Unlock()

			n.logger.Println("tick | height:", height, "| state:", state)
		}
	}
}

func (n *Node) Stop() error {
	n.mu.Lock()
	if n.state == StateStopping {
		n.mu.Unlock()
		return nil
	}
	n.state = StateStopping
	n.mu.Unlock()

	n.logger.Println("state:", StateStopping)

	n.cancel()

	// stop health first (fast)
	{
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_ = n.health.Stop(ctx)
}

	// stop network
	if err := n.network.Stop(); err != nil {
		n.logger.Println("network stop error:", err)
	}

	n.logger.Println("node stopped")
	return nil
}
