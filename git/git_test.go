package git_test

import (
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/msound/tagger/git"
	"github.com/stretchr/testify/assert"
)

func TestTagExists(t *testing.T) {
	// Create an empty repo.
	dirname, err := makeEmptyRepo()
	assert.Equal(t, nil, err, "Make empty repo")
	t.Logf("Directory is: %s", dirname)

	// Create and commit README.md.
	err = commitFile(dirname, "README.md", "test\n", "Initial commit")
	assert.Equal(t, nil, err, "Create and commit README.md file")

	// Create a tag.
	err = createTag(dirname, "1.0.0")
	assert.Equal(t, nil, err, "Git tag")

	// Call the function to be tested.
	c := git.Client{dirname}
	result, err := c.TagExists("1.0.0")
	assert.Equal(t, nil, err, "Tag exists")
	assert.Equal(t, true, result, "Tag exists")

	// Teardown the git directory.
	err = os.RemoveAll(dirname)
	assert.Equal(t, nil, err, "Remove directory")
}

func TestTagExistsNotInARepo(t *testing.T) {
	// Create a temp directory.
	dirname, err := makeTempDirectory()
	assert.Equal(t, nil, err, "Make directory")
	t.Logf("Directory is: %s", dirname)

	// Call the function to be tested.
	c := git.Client{dirname}
	result, err := c.TagExists("1.0.0")
	assert.NotEqual(t, nil, err, "Due to non existant repo, err should not be nil")
	assert.Equal(t, false, result, "Tag should not exist as there is no repo")

	// Teardown the git directory.
	err = os.RemoveAll(dirname)
	assert.Equal(t, nil, err, "Remove directory")
}

func TestTagExistsWithoutTag(t *testing.T) {
	// Create am empty repo.
	dirname, err := makeEmptyRepo()
	assert.Equal(t, nil, err, "Make empty repo")
	t.Logf("Directory is: %s", dirname)

	// Create and commit README.md.
	err = commitFile(dirname, "README.md", "test\n", "Initial commit")
	assert.Equal(t, nil, err, "Create and commit README.md file")

	// Call the function to be tested.
	c := git.Client{dirname}
	result, err := c.TagExists("1.0.0")
	assert.NotEqual(t, nil, err, "Due to non existant tag, err should not be nil")
	assert.Equal(t, false, result, "Tag should not exist as we did not create it")

	// Teardown the git directory.
	err = os.RemoveAll(dirname)
	assert.Equal(t, nil, err, "Remove directory")
}

func TestChangelog(t *testing.T) {
	// Create an empty repo.
	dirname, err := makeEmptyRepo()
	assert.Equal(t, nil, err, "Make empty repo")
	t.Logf("Directory is: %s", dirname)

	// Create and commit README.md.
	err = commitFile(dirname, "README.md", "test\n", "Initial commit")
	assert.Equal(t, nil, err, "Create and commit README.md file")

	// Create a tag.
	err = createAnnotatedTag(dirname, "1.0.0", "Tagging")
	assert.Equal(t, nil, err, "Git tag")

	// Create and commit another file.
	err = commitFile(dirname, "SECOND.md", "test\n", "Merge pull request #2 from dev/branch\n\nAdded SECOND.md")
	assert.Equal(t, nil, err, "Create and commit SECOND.md file")

	// Call the function to be tested.
	c := git.Client{dirname}
	result, err := c.Changelog("1.0.0")
	assert.Nil(t, err, "Getting Changelog")
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "#2 Added SECOND.md", result[0])

	// Teardown the git directory.
	err = os.RemoveAll(dirname)
	assert.Equal(t, nil, err, "Remove directory")
}

func TestChangelogNotInARepo(t *testing.T) {
	// Create a temp directory.
	dirname, err := makeTempDirectory()
	assert.Equal(t, nil, err, "Make directory")
	t.Logf("Directory is: %s", dirname)

	// Call the function to be tested.
	c := git.Client{dirname}
	result, err := c.Changelog("1.0.0")
	assert.NotEqual(t, nil, err, "Due to non existant repo, err should not be nil")
	assert.Equal(t, 0, len(result), "Changelog should be empty as there is no repo")

	// Teardown the git directory.
	err = os.RemoveAll(dirname)
	assert.Equal(t, nil, err, "Remove directory")
}

func TestChangelogWithoutTag(t *testing.T) {
	// Create am empty repo.
	dirname, err := makeEmptyRepo()
	assert.Equal(t, nil, err, "Make empty repo")
	t.Logf("Directory is: %s", dirname)

	// Create and commit README.md.
	err = commitFile(dirname, "README.md", "test\n", "Initial commit")
	assert.Equal(t, nil, err, "Create and commit README.md file")

	// Call the function to be tested.
	c := git.Client{dirname}
	result, err := c.Changelog("1.0.0")
	assert.NotEqual(t, nil, err, "Due to non existant tag, err should not be nil")
	assert.Equal(t, 0, len(result), "Changelog should be empty as there is no tag")

	// Teardown the git directory.
	err = os.RemoveAll(dirname)
	assert.Equal(t, nil, err, "Remove directory")
}

func TestChangelogWithLightweightTag(t *testing.T) {
	// Create an empty repo.
	dirname, err := makeEmptyRepo()
	assert.Equal(t, nil, err, "Make empty repo")
	t.Logf("Directory is: %s", dirname)

	// Create and commit README.md.
	err = commitFile(dirname, "README.md", "test\n", "Initial commit")
	assert.Equal(t, nil, err, "Create and commit README.md file")

	// Create a tag.
	err = createTag(dirname, "1.0.0")
	assert.Equal(t, nil, err, "Git tag")

	// Call the function to be tested.
	c := git.Client{dirname}
	result, err := c.Changelog("1.0.0")
	assert.NotNil(t, err, "Due to lightweight tag, err should not be nil")
	assert.Equal(t, 0, len(result), "Changelog shoud be empty as there is no annotated tag")

	// Teardown the git directory.
	err = os.RemoveAll(dirname)
	assert.Equal(t, nil, err, "Remove directory")
}

func randomString() string {
	value := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		c := rand.Intn(26)
		value = value + string(97+c)
	}

	return value
}

func makeTempDirectory() (string, error) {
	dirname := filepath.Join(os.TempDir(), randomString())
	err := os.Mkdir(dirname, 0777)
	return dirname, err
}

func makeEmptyRepo() (string, error) {
	// Create a temp directory.
	dirname, err := makeTempDirectory()
	if err != nil {
		return dirname, err
	}

	// Initialize git in the temp directory.
	cmd1 := exec.Command("git", "init")
	cmd1.Dir = dirname
	err = cmd1.Run()

	return dirname, err
}

func makeFile(dirname string, filename string, contents string) error {
	file := filepath.Join(dirname, filename)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(contents))
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func commitFile(dirname, filename, contents, message string) error {
	// Create file.
	err := makeFile(dirname, filename, contents)
	if err != nil {
		return err
	}

	// Add file to git staging area.
	cmd2 := exec.Command("git", "add", filename)
	cmd2.Dir = dirname
	err = cmd2.Run()
	if err != nil {
		return err
	}

	// Commit the file.
	cmd3 := exec.Command("git", "commit", "-m", message)
	cmd3.Dir = dirname
	err = cmd3.Run()

	return err
}

func createTag(dirname, tag string) error {
	cmd := exec.Command("git", "tag", tag)
	cmd.Dir = dirname
	err := cmd.Run()

	return err
}

func createAnnotatedTag(dirname, tag, annotation string) error {
	cmd := exec.Command("git", "tag", "-a", tag, "-m", annotation)
	cmd.Dir = dirname
	err := cmd.Run()

	return err
}
