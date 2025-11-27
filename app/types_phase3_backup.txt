package app

// AppType contains high-level metadata for the NOORCHAIN application.
// In Phase 2 this remains minimal.
type AppType struct {
	Name    string
	Version string
}

// DefaultAppType returns the default NOORCHAIN metadata.
func DefaultAppType() AppType {
	return AppType{
		Name:    "NOORCHAIN",
		Version: "0.1.0",
	}
}
