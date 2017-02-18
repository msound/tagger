// This file contains helper functions for loading config from file.
package main

import (
	"errors"

	yaml "gopkg.in/yaml.v2"
)

// LoadConfig loads config from .tagger.yml file.
func LoadConfig(config *Config) error {
	data, err := GetFileContents(".tagger.yml")
	if err != nil {
		return errors.New("Cannot open .tagger.yml")
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return errors.New("Cannot load config from .tagger.yml")
	}

	return nil
}
