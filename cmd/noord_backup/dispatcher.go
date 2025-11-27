package main

// Dispatch executes the default NOORCHAIN command flow.
// It now uses the updated Start() logic based on StartNOORChain().
func Dispatch() error {
	cmds := NewCommands()
	return cmds.Start()
}
