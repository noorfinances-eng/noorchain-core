package app

// AppContext groups minimal application-wide resources.
// In Phase 2 this remains empty and acts as a structural placeholder.
type AppContext struct{}

// NewAppContext returns an empty AppContext placeholder.
func NewAppContext() AppContext {
	return AppContext{}
}
