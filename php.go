package main

import (
	"errors"
	"regexp"

	semver "github.com/coreos/go-semver/semver"
)

// PHPGetSemver extracts semantic version from a PHP file.
func PHPGetSemver() (*semver.Version, error) {
	// open version file
	rawdata, err := GetFileContents(config.FilePath)
	if err != nil {
		return nil, errors.New("Error opening version file")
	}

	// extract the version
	contents := string(rawdata)
	re := regexp.MustCompile(config.VersionKey + ".*'(.*)'")
	parts := re.FindStringSubmatch(contents)

	if len(parts) < 2 {
		return nil, errors.New("Error extracting version from version file")
	}

	version := parts[1]
	return semver.New(version), nil
}
