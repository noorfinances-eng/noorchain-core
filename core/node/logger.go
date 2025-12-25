package node

import (
	"log"
	"os"
)

func newLogger() Logger {
	// Stable prefix + timestamps; stdout
	return log.New(os.Stdout, "[node] ", log.LstdFlags)
}
