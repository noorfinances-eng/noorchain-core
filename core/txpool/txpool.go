package txpool

import "sync"

// Tx is a minimal container for a signed raw transaction (RLP bytes) tracked by hash.
type Tx struct {
	Hash  string // 0x...
	Raw   []byte // signed raw tx bytes (RLP)
	From  string // optional (M8.A)
	To    string // optional (M8.A)
	Nonce uint64 // optional (M8.A)
}

type Pool struct {
	mu      sync.RWMutex
	pending []Tx
	byHash  map[string]Tx
}

func New() *Pool {
	return &Pool{
		byHash: make(map[string]Tx),
	}
}

// AddPending stores tx as pending.
func (p *Pool) AddPending(tx Tx) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.pending = append(p.pending, tx)
	p.byHash[tx.Hash] = tx
}

// PopPending pops up to max pending txs (FIFO).
func (p *Pool) PopPending(max int) []Tx {
	p.mu.Lock()
	defer p.mu.Unlock()

	if max <= 0 || len(p.pending) == 0 {
		return nil
	}
	if max > len(p.pending) {
		max = len(p.pending)
	}
	out := make([]Tx, max)
	copy(out, p.pending[:max])
	p.pending = p.pending[max:]
	return out
}

// Get returns tx by hash if present.
func (p *Pool) Get(hash string) (Tx, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	tx, ok := p.byHash[hash]
	return tx, ok
}
