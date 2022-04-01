package services

import (
	"goteams/internal/repository"
)

func Fetch() ([]Player, error) {
	repo, err := repository.New()
	if err != nil {
		return nil, err
	}

	var players []Player

	if err := repo.Find(&players).Error; err != nil {
		return nil, err
	}

	return players, nil
}

func FetchById(id string) (*Player, error) {
	repo, err := repository.New()
	if err != nil {
		return nil, err
	}

	var player Player

	if err := repo.First(&player, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &player, nil
}
