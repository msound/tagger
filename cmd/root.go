package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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

var config Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tagger",
	Short: "Tagger helps bump version numbers in git projects",
	Long: `Tagger helps developers to bump version numbers in git projects.
Tagger also generates changelog based on merge commits.
Additionally, Tagger creates annotated tags with changelog information.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Search config in current directory with name ".tagger" (without extension).
	viper.AddConfigPath(".")
	viper.SetConfigName(".tagger")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		config.DefaultBranch = viper.GetString("branch")
		config.UpstreamRemote = viper.GetString("remote")
		config.FileFormat = viper.GetString("format")
		config.FilePath = viper.GetString("file")
		config.VersionKey = viper.GetString("key")
	}
}
