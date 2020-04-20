package cmd

import (
	"github.com/spf13/cobra"
	"github.com/theNullP0inter/kenobi/app"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Executes custom commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

type ExecCommand func(appInstance *app.App) *cobra.Command

func init() {

	registeredCommands, ok := appInstance.Config["CommandRegistry"]

	if ok {
		for _, command := range registeredCommands.([]ExecCommand) {
			execCmd.AddCommand(command(appInstance))
		}
	}

	rootCmd.AddCommand(execCmd)
}
