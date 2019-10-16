package git

import (
	"errors"
	"io"
	"regexp"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

// Client has helpful functions to help deal with git.
type Client struct {
	Path string
}

// TagExists confirms if a given tag exists in the git repo.
func (c *Client) TagExists(tag string) (bool, error) {
	repo, err := git.PlainOpen(c.Path)
	if err != nil {
		return false, err
	}

	_, err = repo.Tag(tag)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Changelog returns an array of changes since the given tag.
func (c *Client) Changelog(sinceTag string) ([]string, error) {
	repo, err := git.PlainOpen(c.Path)
	if err != nil {
		return nil, err
	}

	tag, err := repo.Tag(sinceTag)
	if err != nil {
		return nil, err
	}

	tagobj, err := repo.TagObject(tag.Hash())
	if err != nil {
		return nil, err
	}

	commit, err := tagobj.Commit()
	if err != nil {
		return nil, err
	}

	cIter, err := repo.Log(&git.LogOptions{Order: git.LogOrderCommitterTime})
	if err != nil {
		return nil, err
	}

	var result []string

	foundTagCommit := false
	for c, err := cIter.Next(); err != io.EOF; c, err = cIter.Next() {
		if c.Hash == commit.Hash {
			foundTagCommit = true
			break
		}
		if isMergeCommit(c.Message) {
			change := formatCommitMessage(strings.TrimSpace(c.Message))
			result = append(result, change)
		}

	}

	if !foundTagCommit {
		return nil, errors.New("Tag not found in this branch")
	}

	return result, nil
}

func isMergeCommit(msg string) bool {
	return strings.HasPrefix(strings.TrimSpace(msg), "Merge pull request")
}

func formatCommitMessage(msg string) string {
	re := regexp.MustCompile("Merge pull request ([#0-9]*)")
	matches := re.FindSubmatch([]byte(msg))
	if matches == nil {
		return ""
	}

	pullRequestNumber := string(matches[1])

	lines := strings.Split(msg, "\n")
	if len(lines) < 2 {
		return ""
	}

	commitMessage := lines[2]

	return pullRequestNumber + " " + commitMessage
}
