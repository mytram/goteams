package services

import (
	"strings"

	repo "goteams/internal/repository/teams"
)

func Create(dto *CreateTeam) (*Team, error) {
	team := createToTeam(dto)
	// TODO: Add validation

	return repo.Create(team)
}

func createToTeam(dto *CreateTeam) *Team {
	return &Team{
		Name: strings.TrimSpace(dto.Name),
	}
}
