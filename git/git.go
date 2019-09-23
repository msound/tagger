package git

import (
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
