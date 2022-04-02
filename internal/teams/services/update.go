package services

import (
	"strings"

	"gorm.io/gorm"

	"goteams/internal/repository"
)

func Update(teamID string, dto *UpdateTeam) (*Team, error) {
	repo, err := repository.New()
	if err != nil {
		return nil, err
	}

	team := &Team{}

	if err := repo.First(&team, "id = ?", teamID).Error; err != nil {
		return nil, err
	}

	if err := repo.Transaction(updateTeamTx(team, dto)); err != nil {
		return nil, err
	}

	return team, nil
}

func updateTeamTx(team *Team, dto *UpdateTeam) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		if err = updatePlayers(tx, team, dto); err != nil {
			return
		}

		if name := strings.TrimSpace(dto.Name); name != "" {
			team.Name = name
			err = tx.Save(team).Error
		}

		return
	}
}

func updatePlayers(tx *gorm.DB, team *Team, dto *UpdateTeam) error {
	if dto.PlayerIDs == nil {
		return nil
	}

	var players []*Player

	if len(dto.PlayerIDs) > 0 {
		if err := tx.Find(&players, dto.PlayerIDs).Error; err != nil {
			return err
		}
	}

	return tx.Model(team).Association("Players").Replace(players)
}
