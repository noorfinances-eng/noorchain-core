package network

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type Network struct {
	addr string
	ln   net.Listener

	mu    sync.Mutex
	peers map[string]net.Conn
	done  chan struct{}
}

func New(addr string) *Network {
	return &Network{
		addr:  addr,
		peers: make(map[string]net.Conn),
		done:  make(chan struct{}),
	}
}

func (n *Network) Start() error {
	ln, err := net.Listen("tcp", n.addr)
	if err != nil {
		return fmt.Errorf("network listen %s: %w", n.addr, err)
	}
	n.ln = ln

	log.New(os.Stdout, "[network] ", log.LstdFlags).Println("listening on", n.addr)

	// Minimal accept loop: garde les connexions entrantes ouvertes
	go n.acceptLoop()

	return nil
}

func (n *Network) acceptLoop() {
	logger := log.New(os.Stdout, "[p2p] ", log.LstdFlags)

	for {
		conn, err := n.ln.Accept()
		if err != nil {
			select {
			case <-n.done:
				return
			default:
				logger.Println("accept error:", err)
				continue
			}
		}

		remote := conn.RemoteAddr().String()

		n.mu.Lock()
		n.peers[remote] = conn
		cnt := len(n.peers)
		n.mu.Unlock()

		logger.Printf("inbound peer registered %s (peers=%d)", remote, cnt)

		// On ne lit/écrit pas encore de payload (MVP connectivité).
		_ = conn.SetDeadline(time.Now().Add(24 * time.Hour))
	}
}

// Connect dials un peer et garde la session ouverte sous gestion du Network.
func (n *Network) Connect(addr string) error {
	addr = strings.TrimSpace(addr)
	if addr == "" {
		return nil
	}

	logger := log.New(os.Stdout, "[p2p] ", log.LstdFlags)
	logger.Printf("dialing peer %s", addr)

	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		logger.Printf("failed to connect to peer %s: %v", addr, err)
		return err
	}

	remote := conn.RemoteAddr().String()

	n.mu.Lock()
	n.peers[remote] = conn
	cnt := len(n.peers)
	n.mu.Unlock()

	logger.Printf("peer registered %s (peers=%d)", remote, cnt)

	_ = conn.SetDeadline(time.Now().Add(24 * time.Hour))
	return nil
}

func (n *Network) Stop() error {
	close(n.done)

	n.mu.Lock()
	for _, c := range n.peers {
		_ = c.Close()
	}
	n.peers = make(map[string]net.Conn)
	n.mu.Unlock()

	if n.ln != nil {
		if err := n.ln.Close(); err != nil {
			log.New(os.Stdout, "[network] ", log.LstdFlags).Println("close error:", err)
			return err
		}
	}

	log.New(os.Stdout, "[network] ", log.LstdFlags).Println("stopped")
	return nil
}
