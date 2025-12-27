package node

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/syndtr/goleveldb/leveldb"

	"noorchain-evm-l1/core/config"
	"noorchain-evm-l1/core/health"
	"noorchain-evm-l1/core/network"
	"noorchain-evm-l1/core/txindex"
	"noorchain-evm-l1/core/txpool"
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

	txpool  *txpool.Pool
	txindex *txindex.Index

	db *leveldb.DB

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
		txpool:  txpool.New(),
		txindex: txindex.New(),
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

	db, err := openLevelDB(n.cfg.DataDir)
	if err != nil {
		return err
	}
	n.db = db

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

			// M8.A minimal inclusion: mark some pending txs as mined at this height
			mined := 0
			if n.txpool != nil && n.txindex != nil {
				txs := n.txpool.PopPending(64)
				for i := range txs {
					n.txindex.Put(txs[i].Hash, height)
					mined++
				}
			}

			if mined > 0 {
				n.logger.Println("tick | height:", height, "| state:", state, "| mined:", mined)
			} else {
				n.logger.Println("tick | height:", height, "| state:", state)
			}
		}
	}
}

func (n *Node) DB() *leveldb.DB { return n.db }

func (n *Node) Height() uint64 {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.height
}

func (n *Node) TxPool() *txpool.Pool    { return n.txpool }
func (n *Node) TxIndex() *txindex.Index { return n.txindex }

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

	if n.db != nil {
		_ = n.db.Close()
		n.db = nil
	}

	n.logger.Println("node stopped")
	return nil
}
