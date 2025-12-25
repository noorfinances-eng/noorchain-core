package network

import (
	"fmt"
	"net"
)

type Network struct {
	addr string
	ln   net.Listener
}

func New(addr string) *Network {
	return &Network{addr: addr}
}

func (n *Network) Start() {
	ln, err := net.Listen("tcp", n.addr)
	if err != nil {
		panic(err)
	}
	n.ln = ln
	fmt.Println("network listening on", n.addr)
}

func (n *Network) Stop() {
	if n.ln != nil {
		_ = n.ln.Close()
	}
	fmt.Println("network stopped")
}
