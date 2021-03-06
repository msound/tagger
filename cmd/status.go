package cmd

import (
	"fmt"

	"github.com/msound/tagger/config"
	"github.com/msound/tagger/git"
	"github.com/msound/tagger/version"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "tagger status displays last version",
	Long: `tagger status displays last version.
It also displays how many commits have been made since the last tag.
`,
	Run: getStatus,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func getStatus(cmd *cobra.Command, args []string) {
	// TODO : Check if user is on correct branch.
	// Read in version from file and display.
	vrw, err := version.MakeVersionReaderWriter(config.C)
	if err != nil {
		fmt.Println("Error reading version: ", err)
		return
	}

	v, err := vrw.ReadVersion()
	if err != nil {
		fmt.Println("Error reading version: ", err)
		return
	}
	fmt.Println("Current Version is: " + v)

	g := git.Client{Path: "."}
	result, err := g.TagExists(v)
	if err != nil {
		fmt.Println("Error checking if tag exists: ", err)
	}
	if !result {
		fmt.Println("Tag cannot be found. Did you run `git fetch <remote>`")
	}

	changelog, err := g.Changelog(v)
	if err != nil {
		fmt.Print("Error getting changelog: ", err)
	}
	fmt.Println("*** CHANGELOG ***")
	for _, line := range changelog {
		fmt.Println(line)
	}
}
