package services

import (
	"goteams/internal/repository"
)

func Fetch() (teams []Team, err error) {
	repo, err := repository.New()
	if err != nil {
		return
	}

	result := repo.Preload("Players").Find(&teams)
	if result.Error != nil {
		return nil, result.Error
	}

	return
}

func FetchById(id string) (team *Team, err error) {
	repo, err := repository.New()
	if err != nil {
		return
	}

	result := repo.Preload("Players").First(&team, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return
}
