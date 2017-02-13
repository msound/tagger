package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	git2go "gopkg.in/libgit2/git2go.v25"
)

// Flags is a struct that contains various command line flags.
type Flags struct {
	format  string // Format of the version file
	verbose bool
}

// Config is a struct that contains default configuration.
type Config struct {
	defaultBranch  string
	upstreamRemote string
}

var flags = new(Flags)
var config = new(Config)

func init() {
	flag.StringVar(&flags.format, "format", "php", "Format of the version file")
	flag.BoolVar(&flags.verbose, "verbose", false, "Verbose output")
	flag.BoolVar(&flags.verbose, "v", false, "Verbose output (shorthand)")
}

func main() {
	flag.Parse()

	config.defaultBranch = "master"
	config.upstreamRemote = "upstream"

	// Check if we are on master branch.
	repo, err := git2go.OpenRepository("/Users/msound/synbox/synapp")
	FailOnError(err, "Error opening repo")

	name, local, err := GetCurrentBranch(repo)
	FailOnError(err, "Error in getting current branch")

	if name != config.defaultBranch {
		Die("You are not on branch: " + config.defaultBranch)
	}

	// Fetch origin.
	upstream, err := GitFetch(repo, config.upstreamRemote)
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

// GetCurrentBranch returns the current branch name.
func GetCurrentBranch(repo *git2go.Repository) (string, *git2go.Oid, error) {
	ref, err := repo.Head()
	if err != nil {
		return "", nil, errors.New("Cannot read HEAD")
	}
	oid := ref.Target()
	branch := ref.Branch()

	name, err := branch.Name()
	if err != nil {
		return "", nil, errors.New("Cannot get current branch name")
	}

	return name, oid, nil
}

// GitFetch performs a fetch operation on the repo from the given remote.
func GitFetch(repo *git2go.Repository, remoteName string) (*git2go.Oid, error) {
	remote, err := repo.Remotes.Lookup(remoteName)
	if err != nil {
		return nil, errors.New("Cannot find remote: " + remoteName)
	}

	callbacks := git2go.RemoteCallbacks{
		CredentialsCallback:      CredsCallback,
		CertificateCheckCallback: CertCheckCallback,
	}
	proxyopts := git2go.ProxyOptions{}
	var headers []string
	err = remote.ConnectFetch(&callbacks, &proxyopts, headers)
	if err != nil {
		return nil, errors.New("Cannot do git fetch")
	}

	heads, err := remote.Ls("HEAD")
	if err != nil {
		return nil, errors.New("Cannot determine remote HEAD")
	}
	if len(heads) < 1 {
		return nil, errors.New("Cannot determine remote HEAD")
	}

	return heads[0].Id, nil
}

// CredsCallback is a credentials callback function for remote operations.
func CredsCallback(url string, usernameFromURL string, allowedTypes git2go.CredType) (git2go.ErrorCode, *git2go.Cred) {
	ret, cred := git2go.NewCredSshKeyFromAgent("git")
	return git2go.ErrorCode(ret), &cred
}

// CertCheckCallback is a callback function to validate certificate.
func CertCheckCallback(cert *git2go.Certificate, valid bool, hostname string) git2go.ErrorCode {
	if cert.Kind == git2go.CertificateHostkey {
		return git2go.ErrorCode(git2go.ErrOk)
	}

	Die("Tagger does not support HTTPS for git remote.")
	return git2go.ErrorCode(git2go.ErrGeneric)
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
