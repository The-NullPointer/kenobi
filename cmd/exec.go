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

type CommandCenter []func(*app.App) *cobra.Command
