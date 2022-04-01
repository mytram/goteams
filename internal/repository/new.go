package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	*gorm.DB
}

var repository *Repository
var dbFile string

func Setup(file string) {
	dbFile = file
}

func New() (*Repository, error) {
	if repository != nil {
		return repository, nil
	}

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{DB: db}, nil
}
