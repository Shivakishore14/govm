package cmd

import (
	"fmt"
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Display the versions installed",
	Long:  `Display all the versions of Go installed`,
	Run: func(cmd *cobra.Command, args []string) {
		remoteVersions := engine.LocalList()
		//fmt.Println(remoteVersions)
		for _, x := range remoteVersions {
			fmt.Println(x.Name, x.DownloadLink)
		}
	},
}
