package app

import (
	"fmt"
	"github.com/theNullP0inter/kenobi/config"
	"github.com/theNullP0inter/kenobi/db"
	"github.com/theNullP0inter/kenobi/defaults"
)

type App struct {
	Config   config.Config
	Database *db.Database
}

func (a *App) validateConfig() (err error) {

	conf := a.Config

	key, ok := conf["SecretKey"]

	if !ok || len(key.(string)) == 0 {
		return fmt.Errorf("SecretKey must be set")
	}

	return nil

}

func (a *App) InitDb() (err error) {

	conf := a.Config

	dialect, ok := conf["DatabaseDialect"]

	if !ok {
		return fmt.Errorf("DatabaseDialect must be set")
	}

	uri, ok := conf["DatabaseUri"]

	if !ok {
		return fmt.Errorf("DatabaseUri must be set")
	}

	db, err := db.New(dialect.(string), uri.(string))

	if err != nil {
		return err
	}

	a.Database = db

	return nil

}

func (a *App) Close() error {
	return a.Database.Close()
}

func New(conf config.Config) (app *App, err error) {
	appConf := defaults.Config

	for k, v := range conf {
		appConf[k] = v
	}

	app = &App{Config: appConf}
	err = app.validateConfig()

	if err != nil {
		return app, err
	}

	return app, nil

}
