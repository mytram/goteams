package services

import (
	"goteams/internal/repository"
)

func Fetch() ([]Team, error) {
	repo, err := repository.New()
	if err != nil {
		return nil, err
	}

	var teams []Team

	if err := repo.Preload("Players").Find(&teams).Error; err != nil {
		return nil, err
	}

	return teams, nil
}

func FetchById(id string) (*Team, error) {
	repo, err := repository.New()
	if err != nil {
		return nil, err
	}

	var team Team

	if err := repo.Preload("Players").First(&team, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &team, nil
}
