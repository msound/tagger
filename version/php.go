package version

import (
	"errors"
	"os"
	"regexp"

	"github.com/msound/tagger/config"
)

// PHPVersion reads/writes version for a .php file
type PHPVersion struct {
	c config.Config
}

// ReadVersion from a .php file
func (p *PHPVersion) ReadVersion() (string, error) {
	var buf []byte
	buf = make([]byte, 1024)

	file, err := os.Open(p.c.FilePath)
	if err != nil {
		return "", err
	}

	_, err = file.Read(buf)
	if err != nil {
		return "", err
	}

	return parseVersion(buf, p.c.VersionKey)
}

func parseVersion(buf []byte, versionKey string) (string, error) {
	re := regexp.MustCompile("define\\s*\\(\\s*[\\'\\\"]" + versionKey + "[\\'\\\"]\\s*,\\s*[\\'\\\"](.*)[\\'\\\"]\\s*\\)")
	matches := re.FindSubmatch(buf)
	if matches == nil {
		return "", errors.New("Cannot parse version")
	}

	return string(matches[1]), nil
}

// WriteVersion to a .php file
func (p *PHPVersion) WriteVersion(version string) error {
	// Do nothing
	return nil
}
