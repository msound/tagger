package version

import "github.com/msound/tagger/config"

// ReadWriter interface supports methods to read and write versions.
type ReadWriter interface {
	ReadVersion(filePath string, versionKey string) string
	WriteVersion(filePath string, versionKey string, version string)
}

// Manager manipulates versions in files
type Manager struct {
	config config.Config
	ReadWriter
}

// MakeManager is a factory method to create the variable
func MakeManager(config config.Config) *Manager {
	var v *Manager
	switch config.FileFormat {
	case "php":
		v = &Manager{
			config,
			PHPVersion{},
		}
	case "yaml":
		v = &Manager{
			config,
			YAMLVersion{},
		}
	}

	return v
}
