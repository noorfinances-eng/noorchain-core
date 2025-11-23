package app

// Logger is a placeholder for future NOORCHAIN logging capabilities.
// In Phase 2 it remains empty. It will be wired to a real logging backend later.
type Logger struct{}

// NewLogger returns an empty Logger placeholder.
func NewLogger() Logger {
	return Logger{}
}
