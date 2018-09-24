package cmd

import (
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "execute with version",
	Long: `execute the command with the version specified
			$ govm exec 1.10 go ...
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// it should have version number and an extra arg
		if len(args) < 2 {
			log.Println("Not valid format \nUsage: `govm exec 1.10 go <cmd>`")
			return
		}
		versionName := args[0]
		if !strings.HasPrefix(versionName, "go") {
			versionName = "go" + versionName
		}
		if err := engine.Exec(versionName, args[1:]); err != nil {
			log.Println(err)
		}
	},
}
