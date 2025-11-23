package main

// Dispatch runs the default NOORCHAIN command flow.
// In Phase 2 it simply calls the "start" command via Commands.
func Dispatch() error {
	cmds := NewCommands()
	return cmds.Start()
}
