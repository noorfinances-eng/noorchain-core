package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"noorchain-evm-l1/core/config"
	"noorchain-evm-l1/core/network"
)

func main() {
	mainlog := log.New(os.Stdout, "[main] ", log.LstdFlags)
	mainlog.Println("NOORCHAIN 2.1 — EVM L1 booting")

	chainID := flag.String("chain-id", "", "chain identifier")
	dataDir := flag.String("data-dir", "", "data directory")
	p2pAddr := flag.String("p2p-addr", "", "p2p listen address")
	bootPeers := flag.String("boot-peers", "", "comma-separated list of boot peers (host:port)")
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

	// Initialisation réseau DIRECTE (pas de Node abstrait)
	netw := network.New(cfg.P2PAddr)
	if err := netw.Start(); err != nil {
		mainlog.Println("fatal: network start failed:", err)
		os.Exit(1)
	}

	// Bootstrap P2P best-effort
	if peers := strings.TrimSpace(*bootPeers); peers != "" {
		for _, raw := range strings.Split(peers, ",") {
			addr := strings.TrimSpace(raw)
			if addr == "" {
				continue
			}
			_ = netw.Connect(addr)
		}
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	if err := netw.Stop(); err != nil {
		mainlog.Println("warn: network stop error:", err)
	}
	mainlog.Println("NOORCHAIN 2.1 — shutdown clean")
}
