package cmd

import (
	"fmt"

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
	fmt.Println("status called")
	fmt.Println("branch : " + config.DefaultBranch)
	// Check if user is on correct branch.
	// Read in version from file.
	// Display last version.
}
