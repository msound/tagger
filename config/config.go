package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is a struct that contains default configuration.
type Config struct {
	DefaultBranch  string `yaml:"branch"` // Default branch of the repo. ex: master
	UpstreamRemote string `yaml:"remote"` // Upstream remote of the repo. ex: origin
	FileFormat     string `yaml:"format"` // Format of file that contains version. ex: php, yaml, json
	FilePath       string `yaml:"file"`   // Path to file that contains version. ex: docroot/version.php
	VersionKey     string `yaml:"key"`    // Key in the file that refers to the version. ex: APP_VERSION
}

// C contains the configuration read in from the config file.
var C Config

// initConfig reads in config file and ENV variables if set.
func init() {
	// Search config in current directory with name ".tagger" (without extension).
	viper.AddConfigPath(".")
	viper.SetConfigName(".tagger")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		C.DefaultBranch = viper.GetString("branch")
		C.UpstreamRemote = viper.GetString("remote")
		C.FileFormat = viper.GetString("format")
		C.FilePath = viper.GetString("file")
		C.VersionKey = viper.GetString("key")
	}
}
