package version

import (
	"errors"

	"github.com/msound/tagger/config"
)

// ReaderWriter is a struct that has methods to read and write versions.
type ReaderWriter interface {
	ReadVersion() (string, error)
	WriteVersion(version string) error
}

// MakeVersionReaderWriter is a factory method to create the variable
func MakeVersionReaderWriter(c config.Config) (ReaderWriter, error) {

	switch c.FileFormat {
	case "php":
		return &PHPVersion{c}, nil
	}

	return nil, errors.New("The given file format is not supported")
}
