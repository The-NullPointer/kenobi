package app

import (
	"fmt"
	"kenobi/config"
)

type App struct {
	Config config.Config
}

func (a *App) validateConfig() (err error) {

	conf := a.Config

	key, ok := conf["SecretKey"]

	if !ok || len(key.(string)) == 0 {
		return fmt.Errorf("SecretKey must be set")
	}

	return nil

}

func New(conf config.Config) (app *App, err error) {
	app = &App{Config: conf}
	err = app.validateConfig()

	if err != nil {
		return app, err
	}

}
