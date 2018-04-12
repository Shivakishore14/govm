package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Use command is to be called from script,
// this serves as a warning for users and to display command info.
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Lets you switch to another go version",
	Long:  `Use command lets you switch between multiple go versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please check configuration \n run `govm configure`")
	},
}
