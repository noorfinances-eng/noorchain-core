package txindex

import "sync"

// Index maps tx hash -> block number (M8.A: in-memory, non-persistent).
type Index struct {
	mu sync.RWMutex
	m  map[string]uint64
}

func New() *Index {
	return &Index{m: make(map[string]uint64)}
}

func (i *Index) Put(hash string, blockNum uint64) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.m[hash] = blockNum
}

func (i *Index) Get(hash string) (uint64, bool) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	v, ok := i.m[hash]
	return v, ok
}
