package network

import (
	"fmt"
	"log"
	"net"
	"os"
)

type Network struct {
	addr string
	ln   net.Listener
}

func New(addr string) *Network {
	return &Network{addr: addr}
}

func (n *Network) Start() error {
	ln, err := net.Listen("tcp", n.addr)
	if err != nil {
		return fmt.Errorf("network listen %s: %w", n.addr, err)
	}
	n.ln = ln

	log.New(os.Stdout, "[network] ", log.LstdFlags).Println("listening on", n.addr)
	return nil
}

func (n *Network) Stop() error {
	if n.ln != nil {
		if err := n.ln.Close(); err != nil {
			// keep stopping even if close fails
			log.New(os.Stdout, "[network] ", log.LstdFlags).Println("close error:", err)
			return err
		}
	}
	log.New(os.Stdout, "[network] ", log.LstdFlags).Println("stopped")
	return nil
}
