package version

// PHPVersion reads/writes version for a .php file
type PHPVersion struct {
}

// ReadVersion from a .php file
func (p PHPVersion) ReadVersion(filePath string, versionKey string) string {
	// Do nothing
	return "1.2.3"
}

// WriteVersion to a .php file
func (p PHPVersion) WriteVersion(filePath string, versionKey string, version string) {
	// Do nothing
}
