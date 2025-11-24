package main

// Dispatch runs the NOORCHAIN command flow.
// For now, the dispatcher always triggers the Start command.
func Dispatch() error {
	cmds := NewCommands()
	return cmds.Start()
}
