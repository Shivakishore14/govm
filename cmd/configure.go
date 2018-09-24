package cmd

import (
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
	"log"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure Govm",
	Long:  `configure Govm with initial data and this is to be run as a part of installation`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := engine.Configure(); err != nil {
			log.Println(err)
		}
	},
}
