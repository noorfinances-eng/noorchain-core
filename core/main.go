package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"noorchain-evm-l1/core/config"
	"noorchain-evm-l1/core/node"
	"noorchain-evm-l1/core/rpc"
)

func main() {
	mainlog := log.New(os.Stdout, "[main] ", log.LstdFlags)
	mainlog.Println("NOORCHAIN 2.1 — EVM L1 booting")

	chainID := flag.String("chain-id", "", "chain identifier")
	dataDir := flag.String("data-dir", "", "data directory")
	p2pAddr := flag.String("p2p-addr", "", "p2p listen address")
	rpcAddr := flag.String("rpc-addr", "", "json-rpc listen address (e.g. 127.0.0.1:8545)")
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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Start mainnet-like node runtime (DB + network + health + tick loop)
	n := node.New(cfg)
	if err := n.Start(); err != nil {
		mainlog.Println("fatal: node start failed:", err)
		os.Exit(1)
	}

	// Start minimal JSON-RPC server if enabled (still backed by LevelDB via node.DB()).
	var rpcSrv *rpc.Server
	if strings.TrimSpace(*rpcAddr) != "" {
		rpcSrv = rpc.New(*rpcAddr, cfg.ChainID, n, n.DB(), log.New(os.Stdout, "[rpc] ", log.LstdFlags))
		if err := rpcSrv.Start(ctx); err != nil {
			mainlog.Println("fatal: rpc start failed:", err)
			_ = n.Stop()
			os.Exit(1)
		}
	}

	<-ctx.Done()

	// Graceful shutdown
	if rpcSrv != nil {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		_ = rpcSrv.Stop(shutdownCtx)
		cancel()
	}
	if err := n.Stop(); err != nil {
		mainlog.Println("warn: node stop error:", err)
	}

	mainlog.Println("NOORCHAIN 2.1 — shutdown clean")
}
