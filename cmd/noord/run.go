package main

// Run is a convenience wrapper around Start.
// In later phases it may handle additional runtime behavior.
func Run() error {
	return Start()
}
