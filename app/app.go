package app

// App is a minimal placeholder for NOORCHAIN Phase 2.
type App struct {
	Info AppInfo
}

// AppInfo holds basic metadata about the app.
type AppInfo struct {
	Name    string
	Version string
}

// NewApp constructs a minimal NOORCHAIN app instance.
func NewApp() *App {
	return &App{
		Info: AppInfo{
			Name:    "NOORCHAIN",
			Version: "0.1.0",
		},
	}
}
