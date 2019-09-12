package version_test

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/msound/tagger/config"
	"github.com/msound/tagger/version"
	"github.com/stretchr/testify/assert"
)

func TestReadVersion(t *testing.T) {
	filename := os.TempDir() + randomString()
	versionKey := "VERSION"
	want := "1.2.3"
	testData := []byte("<?php define('" + versionKey + "', '" + want + "');")
	file, err := os.Create(filename)
	assert.Equal(t, nil, err, "Create file")
	_, err = file.Write(testData)
	assert.Equal(t, nil, err, "Write file")
	err = file.Close()
	assert.Equal(t, nil, err, "Close file")

	testConfig := config.Config{
		DefaultBranch:  "master",
		UpstreamRemote: "origin",
		FileFormat:     "php",
		FilePath:       filename,
		VersionKey:     versionKey,
	}

	p, err := version.MakeVersionReaderWriter(testConfig)
	assert.Equal(t, nil, err, "Factory function")

	got, err := p.ReadVersion()
	assert.Equal(t, nil, err, "Read version")
	assert.Equal(t, want, got)

	err = os.Remove(filename)
	assert.Equal(t, nil, err, "Remove file")
}

func TestReadVersionWrongVersionKey(t *testing.T) {
	filename := os.TempDir() + randomString()
	versionKey := "VERSION"
	want := "1.2.3"
	testData := []byte("<?php define('" + versionKey + "WRONG" + "', '" + want + "');")
	file, err := os.Create(filename)
	assert.Equal(t, nil, err, "Create file")
	_, err = file.Write(testData)
	assert.Equal(t, nil, err, "Write file")
	err = file.Close()
	assert.Equal(t, nil, err, "Close file")

	testConfig := config.Config{
		DefaultBranch:  "master",
		UpstreamRemote: "origin",
		FileFormat:     "php",
		FilePath:       filename,
		VersionKey:     versionKey,
	}

	p, err := version.MakeVersionReaderWriter(testConfig)
	assert.Equal(t, nil, err, "Factory function")

	got, err := p.ReadVersion()
	assert.Equal(t, "Cannot parse version", err.Error(), "Read version")
	assert.Equal(t, "", got, "Due to error version must be empty string")

	err = os.Remove(filename)
	assert.Equal(t, nil, err, "Remove file")
}

func TestReadVersionBadFile(t *testing.T) {
	filename := os.TempDir() + randomString()
	testConfig := config.Config{
		DefaultBranch:  "master",
		UpstreamRemote: "origin",
		FileFormat:     "php",
		FilePath:       filename,
		VersionKey:     "versionKey",
	}

	p, err := version.MakeVersionReaderWriter(testConfig)
	assert.Equal(t, nil, err, "Factory function")

	got, err := p.ReadVersion()
	assert.NotEqual(t, nil, err, "Due to bad file, err should not be nil")
	assert.Equal(t, "", got, "Due to error version must be empty string")
}

func TestReadVersionEmptyFile(t *testing.T) {
	filename := os.TempDir() + randomString()
	file, err := os.Create(filename)
	assert.Equal(t, nil, err, "Create file")
	err = file.Close()
	assert.Equal(t, nil, err, "Close file")

	testConfig := config.Config{
		DefaultBranch:  "master",
		UpstreamRemote: "origin",
		FileFormat:     "php",
		FilePath:       filename,
		VersionKey:     "VERSION",
	}

	p, err := version.MakeVersionReaderWriter(testConfig)
	assert.Equal(t, nil, err, "Factory function")

	got, err := p.ReadVersion()
	assert.NotEqual(t, nil, err, "Read version")
	assert.Equal(t, "", got, "Due to empty file version must be empty string")

	err = os.Remove(filename)
	assert.Equal(t, nil, err, "Remove file")
}

func randomString() string {
	value := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		c := rand.Intn(26)
		value = value + string(97+c)
	}

	return value
}
