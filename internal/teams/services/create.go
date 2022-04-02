package services

import (
	"strings"

	"goteams/internal/repository"
)

func Create(dto *CreateTeam) (*Team, error) {
	repo, err := repository.New()
	if err != nil {
		return nil, err
	}

	team := createToTeam(dto)
	// TODO: Add validation

	if err := repo.DB.Create(team).Error; err != nil {
		return nil, err
	}

	return team, nil
}

func createToTeam(dto *CreateTeam) *Team {
	return &Team{
		Name: strings.TrimSpace(dto.Name),
	}
}
