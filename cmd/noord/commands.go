package main

// Commands groups the high-level NOORCHAIN CLI actions.
// In Phase 2 this is a thin wrapper around placeholder functions.
type Commands struct{}

// NewCommands returns a new Commands collection.
func NewCommands() Commands {
	return Commands{}
}

// Start executes the "start" command.
func (c Commands) Start() error {
	return Start()
}

// Run executes the "run" command.
func (c Commands) Run() error {
	return Run()
}

// Version executes the "version" command.
func (c Commands) Version() string {
	return Version()
}

// Help executes the "help" command.
func (c Commands) Help() string {
	return Help()
}
