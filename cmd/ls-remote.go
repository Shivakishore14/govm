package cmd

import (
	"fmt"
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
	"os"
)

var lsRemoteCmd = &cobra.Command{
	Use:   "ls-remote",
	Short: "Display all the versions available",
	Long:  `Display all the versions of Go available for download`,
	Run: func(cmd *cobra.Command, args []string) {
		hostOs := os.Getenv("GOVMOS")
		hostArch := os.Getenv("GOVMARCH")
		fmt.Println(hostOs, hostArch)

		if hostOs == "" || hostArch == "" {
			fmt.Println("please check configuration \n run `govm configure`")
			return
		}

		remoteVersions := engine.RemoteList(hostOs, hostArch)
		//fmt.Println(remoteVersions)
		for _, x := range remoteVersions {
			fmt.Println(x.Name)
		}
	},
}
