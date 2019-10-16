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
	// Create a temp directory.
	dirname := filepath.Join(os.TempDir(), randomString())
	file := "README.md"
	filename := filepath.Join(dirname, file)
	t.Logf("Directory is: %s", dirname)
	err := os.Mkdir(dirname, 0777)
	assert.Equal(t, nil, err, "Make directory")

	// Create a temp file.
	f, err := os.Create(filename)
	assert.Equal(t, nil, err, "Create README.md file")
	_, err = f.Write([]byte("test\n"))
	assert.Equal(t, nil, err, "Write to README.md file")
	err = f.Close()
	assert.Equal(t, nil, err, "Close README.md file")

	// Initialize git in the temp directory.
	cmd1 := exec.Command("git", "init")
	cmd1.Dir = dirname
	err = cmd1.Run()
	assert.Equal(t, nil, err, "Git init")

	// Add README.md to git staging area.
	cmd2 := exec.Command("git", "add", file)
	cmd2.Dir = dirname
	err = cmd2.Run()
	assert.Equal(t, nil, err, "Git add")

	// Commit the file.
	cmd3 := exec.Command("git", "commit", "-m", "Initial commit")
	cmd3.Dir = dirname
	err = cmd3.Run()
	assert.Equal(t, nil, err, "Git commit")

	// Create a tag.
	cmd4 := exec.Command("git", "tag", "1.0.0")
	cmd4.Dir = dirname
	err = cmd4.Run()
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
	dirname := filepath.Join(os.TempDir(), randomString())
	t.Logf("Directory is: %s", dirname)
	err := os.Mkdir(dirname, 0777)
	assert.Equal(t, nil, err, "Make directory")

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
	// Create a temp directory.
	dirname := filepath.Join(os.TempDir(), randomString())
	file := "README.md"
	filename := filepath.Join(dirname, file)
	t.Logf("Directory is: %s", dirname)
	err := os.Mkdir(dirname, 0777)
	assert.Equal(t, nil, err, "Make directory")

	// Create a temp file.
	f, err := os.Create(filename)
	assert.Equal(t, nil, err, "Create README.md file")
	_, err = f.Write([]byte("test\n"))
	assert.Equal(t, nil, err, "Write to README.md file")
	err = f.Close()
	assert.Equal(t, nil, err, "Close README.md file")

	// Initialize git in the temp directory.
	cmd1 := exec.Command("git", "init")
	cmd1.Dir = dirname
	err = cmd1.Run()
	assert.Equal(t, nil, err, "Git init")

	// Add README.md to git staging area.
	cmd2 := exec.Command("git", "add", file)
	cmd2.Dir = dirname
	err = cmd2.Run()
	assert.Equal(t, nil, err, "Git add")

	// Commit the file.
	cmd3 := exec.Command("git", "commit", "-m", "Initial commit")
	cmd3.Dir = dirname
	err = cmd3.Run()
	assert.Equal(t, nil, err, "Git commit")

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
	// Create a temp directory.
	dirname := filepath.Join(os.TempDir(), randomString())
	file := "README.md"
	filename := filepath.Join(dirname, file)
	t.Logf("Directory is: %s", dirname)
	err := os.Mkdir(dirname, 0777)
	assert.Equal(t, nil, err, "Make directory")

	// Create a temp file.
	f, err := os.Create(filename)
	assert.Equal(t, nil, err, "Create README.md file")
	_, err = f.Write([]byte("test\n"))
	assert.Equal(t, nil, err, "Write to README.md file")
	err = f.Close()
	assert.Equal(t, nil, err, "Close README.md file")

	// Initialize git in the temp directory.
	cmd1 := exec.Command("git", "init")
	cmd1.Dir = dirname
	err = cmd1.Run()
	assert.Equal(t, nil, err, "Git init")

	// Add README.md to git staging area.
	cmd2 := exec.Command("git", "add", file)
	cmd2.Dir = dirname
	err = cmd2.Run()
	assert.Equal(t, nil, err, "Git add README.md")

	// Commit the file.
	cmd3 := exec.Command("git", "commit", "-m", "Initial commit")
	cmd3.Dir = dirname
	err = cmd3.Run()
	assert.Equal(t, nil, err, "Git commit")

	// Create a tag.
	cmd4 := exec.Command("git", "tag", "-a", "1.0.0", "-m", "Tagging")
	cmd4.Dir = dirname
	err = cmd4.Run()
	assert.Equal(t, nil, err, "Git tag")

	// Create another file.
	file2 := "SECOND.md"
	filename2 := filepath.Join(dirname, file2)
	f2, err := os.Create(filename2)
	assert.Equal(t, nil, err, "Create SECOND.md file")
	_, err = f2.Write([]byte("test\n"))
	assert.Equal(t, nil, err, "Write to SECOND.md file")
	err = f2.Close()
	assert.Equal(t, nil, err, "Close SECOND.md file")

	// Add README.md to git staging area.
	cmd5 := exec.Command("git", "add", file2)
	cmd5.Dir = dirname
	err = cmd5.Run()
	assert.Equal(t, nil, err, "Git add SECOND.md")

	// Commit the file.
	cmd6 := exec.Command("git", "commit", "-m", "Merge pull request #2 from dev/branch\n\nAdded SECOND.md")
	cmd6.Dir = dirname
	err = cmd6.Run()
	assert.Equal(t, nil, err, "Git commit")

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

func randomString() string {
	value := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		c := rand.Intn(26)
		value = value + string(97+c)
	}

	return value
}
