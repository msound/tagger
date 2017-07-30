package main

import (
	"flag"
	"fmt"

	"github.com/coreos/go-semver/semver"
)

// Flags is a struct that contains various command line flags.
type Flags struct {
	verbose bool // Verbose output
}

// Config is a struct that contains default configuration.
type Config struct {
	DefaultBranch  string `yaml:"branch"` // Default branch of the repo. ex: master
	UpstreamRemote string `yaml:"remote"` // Upstream remote of the repo. ex: origin
	FileFormat     string `yaml:"format"` // Format of file that contains version. ex: php, yaml, json
	FilePath       string `yaml:"file"`   // Path to file that contains version. ex: docroot/version.php
	VersionKey     string `yaml:"key"`    // Key in the file that refers to the version. ex: APP_VERSION
}

var flags = new(Flags)
var config = new(Config)

func init() {
	flag.BoolVar(&flags.verbose, "v", false, "Verbose output")
}

func main() {
	flag.Parse()

	// Load configuration.
	Verbose("Loading config from .tagger.yml")
	err := LoadConfig(config)
	FailOnError(err, "Error loading config")

	// Check if current branch is master branch.
	Verbose("Opening repo")
	repo, err := OpenRepo(".")
	FailOnError(err, "Error opening repo")

	Verbose("Getting current branch")
	name, localHead, err := GetCurrentBranch(repo)
	FailOnError(err, "Error in getting current branch")

	if name != config.DefaultBranch {
		Die("You are not on branch: " + config.DefaultBranch)
	}

	// Fetch upstream.
	Verbose("Running git fetch")
	upstreamHead, err := GitFetch(repo, config.UpstreamRemote)
	FailOnError(err, "Error doing git fetch")

	// See if local in sync with upstream.
	if localHead.String() != upstreamHead.String() {
		Die("Your local branch is out of sync with upstream")
	}

	// find current version from version file.
	var version *semver.Version
	switch config.FileFormat {
	case "php":
		version, err = PHPGetSemver()
		FailOnError(err, "Error reading version from "+config.FilePath)
	}

	// Iterate over merge commits since that tag and prepare changelog.
	Verbose("Getting changelog since: %s", version.String())
	commits, err := Changelog(repo, version.String())
	FailOnError(err, "Error getting changelog")

	fmt.Printf("Changelog since %s\n", version.String())
	for _, commit := range commits {
		fmt.Println(commit)
	}

	// Interactievely ask user if they want major, minor or patch.
	var tag string
	fmt.Print("Enter tag (major/minor/patch/x.y.z): ")
	fmt.Scanf("%s", &tag)

	// Bump version and update version.php and commit.
	oldVersion := version.String()
	switch tag {
	case "major":
		version.BumpMajor()
	case "minor":
		version.BumpMinor()
	case "patch":
		version.BumpPatch()
	default:
		err = version.Set(tag)
		FailOnError(err, "Invalid version given")
	}

	// write new version to version file.
	switch config.FileFormat {
	case "php":
		err = PHPSetSemver(oldVersion, version)
		FailOnError(err, "Error writing version to "+config.FilePath)
	}

	// commit the version file.
	files := []string{config.FilePath}
	message := "Preparing to tag " + version.String()
	newHead, err := CommitFiles(repo, files, message)
	FailOnError(err, "Error committing changes")
	Verbose("New HEAD is: " + newHead.Id().String())

	// Create annotated Tag.
	Verbose("Tagging %s\n", version.String())
	message = PrepareMessage(commits)
	err = CreateAnnotatedTag(repo, version.String(), newHead, message)
	FailOnError(err, "Error creating tag")

	fmt.Println(version.String())

	// Push master branch with tags.
}
