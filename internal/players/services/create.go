package services

import (
	"strings"

	"goteams/internal/repository"
)

func Create(dto *CreatePlayer) (*Player, error) {
	repo, err := repository.New()
	if err != nil {
		return nil, err
	}

	player := createToPlayer(dto)
	// TODO: Add validation

	if err := repo.DB.Create(player).Error; err != nil {
		return nil, err
	}

	return player, nil
}

func createToPlayer(dto *CreatePlayer) *Player {
	return &Player{
		Name:   strings.TrimSpace(dto.Name),
		Number: dto.Number,
	}
}
