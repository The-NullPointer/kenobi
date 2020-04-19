package app

import (
	"fmt"
	"kenobi/config"
	"kenobi/db"
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

func (a *App) initDb() (err error) {

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
	app = &App{Config: conf}
	err = app.validateConfig()

	if err != nil {
		return app, err
	}

	err = app.initDb()
	if err != nil {
		return app, err
	}

	return app, nil

}
