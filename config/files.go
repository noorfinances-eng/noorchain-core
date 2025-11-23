package config

// File operations placeholder for NOORCHAIN Phase 2.
// In the future this will read/write JSON/TOML/YAML configs.

type File struct {
	Path string
}

// NewFile returns a placeholder File instance.
func NewFile(path string) File {
	return File{Path: path}
}

// Read is a placeholder read function.
func (f File) Read() ([]byte, error) {
	return []byte{}, nil
}

// Write is a placeholder write function.
func (f File) Write(_ []byte) error {
	return nil
}
