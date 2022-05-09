package services

import (
	"strings"

	repo "goteams/internal/repository/players"
)

func Create(dto *CreatePlayer) (*Player, error) {
	player := createToPlayer(dto)
	return repo.Create(player)
}

func createToPlayer(dto *CreatePlayer) *Player {
	return &Player{
		Name:   strings.TrimSpace(dto.Name),
		Number: dto.Number,
	}
}
