package version_test

import (
	"testing"

	"github.com/msound/tagger/config"
	"github.com/msound/tagger/version"
	"github.com/stretchr/testify/assert"
)

func TestMakeVersionReaderWriterWrongFileFormat(t *testing.T) {
	testConfig := config.Config{
		DefaultBranch:  "master",
		UpstreamRemote: "origin",
		FileFormat:     "foo", // This should throw an error.
		FilePath:       "filename",
		VersionKey:     "VERSION",
	}

	vrw, err := version.MakeVersionReaderWriter(testConfig)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, vrw, "Returned value for wrong file format should be nil")
}
