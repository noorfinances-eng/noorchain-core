package node

import (
	"fmt"
	"path/filepath"

	"github.com/syndtr/goleveldb/leveldb"
)

func openLevelDB(dataDir string) (*leveldb.DB, error) {
	path := filepath.Join(dataDir, "db", "leveldb")
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, fmt.Errorf("open leveldb %s: %w", path, err)
	}
	return db, nil
}
