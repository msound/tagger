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

func randomString() string {
	value := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		c := rand.Intn(26)
		value = value + string(97+c)
	}

	return value
}
