package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/theNullP0inter/kenobi/app"
	"github.com/theNullP0inter/kenobi/config"
	"os"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute(conf config.Config, customCommands []ExecCommandFunction) {

	a, err := app.New(conf)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	appInstance = a

	for _, customCommand := range customCommands {
		execCmd.AddCommand(customCommand(appInstance))
	}

	rootCmd.AddCommand(execCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var appInstance *app.App
