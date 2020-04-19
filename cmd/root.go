package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kenobi/config"
	"os"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute(conf *config.Config) {
	appConfig = conf
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var appConfig *config.Config
