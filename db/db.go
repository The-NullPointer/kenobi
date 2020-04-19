package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

type Database struct {
	*gorm.DB
}

func New(dialect string, uri string) (*Database, error) {
	db, err := gorm.Open(dialect, uri)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}
	return &Database{db}, nil
}
