package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"noorchain-evm-l1/core/config"
	"noorchain-evm-l1/core/node"
)

func main() {
	fmt.Println("NOORCHAIN 2.1 — EVM L1 booting")

	cfg := config.Default()
	n := node.New(cfg)
	n.Start()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	n.Stop()
	fmt.Println("NOORCHAIN 2.1 — shutdown clean")
}
