package services

import (
	repo "goteams/internal/repository/teams"
)

func Update(teamID string, dto *UpdateTeam) (*Team, error) {
	return repo.Update(teamID, dto.Name, dto.PlayerIDs)
}
