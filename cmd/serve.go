package cmd

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"kenobi/app"
	"kenobi/router"
	"os"
	"os/signal"
	"sync"
)

func serveApp(ctx context.Context, a *app.App) {

	appRouter := router.New(a)
	appRouter.Init()
	appRouter.Serve(ctx)

}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "runs an Http server",
	RunE: func(cmd *cobra.Command, args []string) error {

		a, err := app.New(appConfig)

		if err != nil {
			return err
		}
		defer a.Close()

		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			<-ch
			logrus.Info("signal caught. shutting down...")
			cancel()
		}()

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer cancel()
			serveApp(ctx, a)
		}()

		wg.Wait()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
