package cmd

import (
	"fmt"
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "finds the path of the go version",
	Long:  `path command finds the path of the go verion given as argument`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Provide version")
			return
		}
		version := args[0]
		if !strings.HasPrefix(version, "go") {
			version = "go" + version
		}
		path, err := engine.Path(version)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("PATH:" +path)
	},
}
