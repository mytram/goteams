package repository

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}

	return &Repository{DB: db}, nil
}
