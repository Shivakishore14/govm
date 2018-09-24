package cmd

import (
	"fmt"
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
	"strings"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a golang version",
	Long:  `Uninstalls Go given the version given`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please pecify a version to uninstall")
			return
		}
		version := args[0]
		if !strings.HasPrefix(version, "go") {
			version = "go" + version
		}
		if err := engine.Uninstall(version); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Successfuly removed ", args[0])

	},
}
