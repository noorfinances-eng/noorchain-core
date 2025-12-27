package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"github.com/syndtr/goleveldb/leveldb"
	"path/filepath"
	"time"

	"noorchain-evm-l1/core/config"
	"noorchain-evm-l1/core/network"
	"noorchain-evm-l1/core/rpc"
)

func main() {
	mainlog := log.New(os.Stdout, "[main] ", log.LstdFlags)
	mainlog.Println("NOORCHAIN 2.1 — EVM L1 booting")

	chainID := flag.String("chain-id", "", "chain identifier")
	dataDir := flag.String("data-dir", "", "data directory")
	p2pAddr := flag.String("p2p-addr", "", "p2p listen address")
	bootPeers := flag.String("boot-peers", "", "comma-separated list of boot peers (host:port)")
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

	// Start P2P network
	netw := network.New(cfg.P2PAddr)
	if err := netw.Start(); err != nil {
		mainlog.Println("fatal: network start failed:", err)
		os.Exit(1)
	}

	// Bootstrap peers (best effort)
	if peers := strings.TrimSpace(*bootPeers); peers != "" {
		for _, raw := range strings.Split(peers, ",") {
			addr := strings.TrimSpace(raw)
			if addr == "" {
				continue
			}
			_ = netw.Connect(addr)
		}
	}
	// Open LevelDB for persistent state (mainnet-like)
	dbPath := filepath.Join(cfg.DataDir, "db", "leveldb")
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		mainlog.Println("fatal: leveldb open failed:", err)
		_ = netw.Stop()
		os.Exit(1)
	}
	defer func() { _ = db.Close() }()


	// Start minimal JSON-RPC server if enabled
	var rpcSrv *rpc.Server
	if strings.TrimSpace(*rpcAddr) != "" {
		rpcSrv = rpc.New(*rpcAddr, cfg.ChainID, db, log.New(os.Stdout, "[rpc] ", log.LstdFlags))
		if err := rpcSrv.Start(ctx); err != nil {
			mainlog.Println("fatal: rpc start failed:", err)
			_ = netw.Stop()
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
	if err := netw.Stop(); err != nil {
		mainlog.Println("warn: network stop error:", err)
	}
	mainlog.Println("NOORCHAIN 2.1 — shutdown clean")
}
