package node

import (
	"log"
	"os"
)

func newLogger() *log.Logger {
	return log.New(os.Stdout, "[node] ", log.LstdFlags)
}
