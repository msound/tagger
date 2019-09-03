package cmd

import (
	"fmt"

	"github.com/msound/tagger/config"
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
	vm := version.MakeManager(config.C)
	v := vm.ReadVersion(config.C.FilePath, config.C.VersionKey)
	fmt.Println("Current Version is: " + v)
}
