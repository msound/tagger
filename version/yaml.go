package version

// YAMLVersion reads/writes version for a .yml file
type YAMLVersion struct {
}

// ReadVersion from a .yml file
func (y YAMLVersion) ReadVersion(filePath string, versionKey string) string {
	// Do nothing
	return "1.2.3"
}

// WriteVersion to a .yml file
func (y YAMLVersion) WriteVersion(filePath string, versionKey string, version string) {
	// Do nothing
}
