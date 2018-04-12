package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "govm",
	Short: "Govm is a version manager for golang",
	Long:  `Govm is a Fast and Flexible Go Version Manager built with go`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(useCmd)
	rootCmd.AddCommand(pathCmd)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(lsRemoteCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(configureCmd)
}
