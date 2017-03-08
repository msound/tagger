package main

import (
	"errors"
	"regexp"
	"strings"

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

// PHPSetSemver writes new version to PHP file.
//
// oldVersion is the old version.
// newVersion contains the new version.
func PHPSetSemver(oldVersion string, newVersion *semver.Version) error {
	// open version file
	rawdata, err := GetFileContents(config.FilePath)
	if err != nil {
		return errors.New("Error opening version file")
	}

	// extract the version
	contents := string(rawdata)
	re := regexp.MustCompile(config.VersionKey + ".*'(.*)'")
	parts := re.FindStringSubmatch(contents)
	find := parts[0]
	replace := strings.Replace(find, oldVersion, newVersion.String(), 1)
	newContents := strings.Replace(contents, find, replace, 1)

	err = PutFileContents(config.FilePath, newContents)
	if err != nil {
		return errors.New("Error writing version file")
	}

	return nil
}
