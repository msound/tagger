// This file contains git related help functions.
package main

import "errors"
import git2go "gopkg.in/libgit2/git2go.v25"

// OpenRepo opens and returns a Repository specified by the path.
func OpenRepo(path string) (*git2go.Repository, error) {
	return git2go.OpenRepository(path)
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
		return nil, errors.New("Error getting remote HEAD")
	}
	if len(heads) < 1 {
		return nil, errors.New("Cannot find remote HEAD")
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
