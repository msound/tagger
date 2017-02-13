package main

import (
	"flag"
	"fmt"
	"os"
)

// Flags is a struct that contains various command line flags.
type Flags struct {
	verbose bool // Verbose output
}

// Config is a struct that contains default configuration.
type Config struct {
	DefaultBranch  string
	UpstreamRemote string
}

var flags = new(Flags)
var config = new(Config)

func init() {
	flag.BoolVar(&flags.verbose, "v", false, "Verbose output")
}

func main() {
	flag.Parse()

	config.DefaultBranch = "master"
	config.UpstreamRemote = "upstream"

	// Check if we are on master branch.
	Verbose("Opening repo")
	repo, err := OpenRepo(".")
	FailOnError(err, "Error opening repo")

	Verbose("Getting current branch")
	name, local, err := GetCurrentBranch(repo)
	FailOnError(err, "Error in getting current branch")

	if name != config.DefaultBranch {
		Die("You are not on branch: " + config.DefaultBranch)
	}

	// Fetch origin.
	Verbose("Running git fetch")
	upstream, err := GitFetch(repo, config.UpstreamRemote)
	FailOnError(err, "Error doing git fetch")

	// See if local in sync with upstream.
	if local.String() != upstream.String() {
		Die("Your local branch is out of sync with upstream")
	}

	fmt.Println("ok")

	// find current version from version.php.

	// Iterate over merge commits since that tag and prepare changelog.

	// Interactievely ask user if they want major, minor or patch.

	// Bump version and update version.php and commit.

	// Create annotated Tag.

	// Push master branch with tags.
}

// Verbose prints debug messages if verbose flag is enabled in command-line options.
func Verbose(format string, a ...interface{}) {
	if flags.verbose {
		fmt.Printf(format+"\n", a...)
	}
}

// FailOnError is a helper function to validate err object and exit if necessary.
func FailOnError(err error, msg string) {
	if err != nil {
		os.Stderr.WriteString(msg + " : " + err.Error() + "\n")
		os.Exit(1)
	}
}

// Die is a helper function to exit the program abnormally.
func Die(format string, a ...interface{}) {
	os.Stderr.WriteString(fmt.Sprintf(format, a...) + "\n")
	os.Exit(1)
}
