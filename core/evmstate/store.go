package evmstate

import (
	"fmt"
	"os"
	"path/filepath"

	gethleveldb "github.com/ethereum/go-ethereum/ethdb/leveldb"
)

// Store owns the geth-compatible key/value database used for EVM world-state
// (trie nodes, code, storage, etc.). It is intentionally isolated from the
// existing NOOR LevelDB (rcpt/v1, poss/v1) to avoid schema collisions.
type Store struct {
	path string
	db   *gethleveldb.Database
}

// Path returns the filesystem path of the geth DB under the given data-dir.
func Path(dataDir string) string {
	return filepath.Join(dataDir, "db", "geth")
}

// Open opens (and creates if needed) the geth-compatible LevelDB at <data-dir>/db/geth.
func Open(dataDir string, readonly bool) (*Store, error) {
	p := Path(dataDir)
	if err := os.MkdirAll(p, 0o755); err != nil {
		return nil, fmt.Errorf("mkdir geth db dir %s: %w", p, err)
	}

	// Conservative defaults (tunable later)
	// cache: 16 MB, handles: 16
	db, err := gethleveldb.New(p, 16, 16, "noorchain-evmstate", readonly)
	if err != nil {
		return nil, fmt.Errorf("open geth leveldb %s: %w", p, err)
	}

	return &Store{path: p, db: db}, nil
}

func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}

func (s *Store) DB() *gethleveldb.Database { return s.db }
func (s *Store) Dir() string               { return s.path }
