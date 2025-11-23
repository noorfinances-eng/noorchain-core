package app

// RuntimeServices is a placeholder grouping for all future NOORCHAIN services:
// RPC server, API server, GRPC, metrics, etc.
// In Phase 2 this structure stays empty.
type RuntimeServices struct{}

// NewRuntimeServices returns an empty services placeholder.
func NewRuntimeServices() RuntimeServices {
	return RuntimeServices{}
}
