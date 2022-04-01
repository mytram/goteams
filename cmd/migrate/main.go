package main

import (
	"os"

	"goteams/internal/models"
	"goteams/internal/repository"
)

func main() {
	os.Setenv("TZ", "UTC")
	repository.Setup(os.Getenv("DB_FILE"))

	repo, err := repository.New()
	if err != nil {
		panic("cannot connect to db: " + err.Error())
	}

	repo.AutoMigrate(
		&models.Team{},
		&models.Player{},
		&models.TeamPlayer{},
	)
}
