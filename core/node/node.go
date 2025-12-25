package node

import "fmt"

type Node struct{}

func New() *Node {
	return &Node{}
}

func (n *Node) Start() {
	fmt.Println("node started")
}

func (n *Node) Stop() {
	fmt.Println("node stopped")
}
