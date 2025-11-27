package main

// Execute runs the NOORCHAIN command dispatcher.
// It now relies on the updated Command flow for the Cosmos-based app.
func Execute() error {
	return Dispatch()
}
