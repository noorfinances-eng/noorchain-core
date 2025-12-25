package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("NOORCHAIN 2.1 — EVM L1 booting")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	fmt.Println("NOORCHAIN 2.1 — shutdown clean")
}
