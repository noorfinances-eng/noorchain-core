package app

import "fmt"

// Coin helpers placeholder for NOORCHAIN Phase 2.

func NewCoin(amount int64) string {
	return fmt.Sprintf("%d%s", amount, DefaultDenom)
}

func NewDisplayCoin(amount int64) string {
	return fmt.Sprintf("%d %s", amount, DefaultDisplayDenom)
}
