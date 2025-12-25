package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"noorchain-evm-l1/core/config"
	"noorchain-evm-l1/core/node"
)

func main() {
	fmt.Println("NOORCHAIN 2.1 — EVM L1 booting")

	chainID := flag.String("chain-id", "", "chain identifier")
	dataDir := flag.String("data-dir", "", "data directory")
	p2pAddr := flag.String("p2p-addr", "", "p2p listen address")
	flag.Parse()

	cfg := config.Default()
	if *chainID != "" {
		cfg.ChainID = *chainID
	}
	if *dataDir != "" {
		cfg.DataDir = *dataDir
	}
	if *p2pAddr != "" {
		cfg.P2PAddr = *p2pAddr
	}

	n := node.New(cfg)
	n.Start()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	n.Stop()
	fmt.Println("NOORCHAIN 2.1 — shutdown clean")
}
