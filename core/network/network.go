package network

import "fmt"

type Network struct {
	addr string
}

func New(addr string) *Network {
	return &Network{addr: addr}
}

func (n *Network) Start() {
	fmt.Println("network listening on", n.addr)
}

func (n *Network) Stop() {
	fmt.Println("network stopped")
}
