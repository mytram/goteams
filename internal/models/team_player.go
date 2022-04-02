package models

type TeamPlayer struct {
	Base
	TeamID   int    `json:"teamId" gorm:"index:idx_team_players_team_player_id,unique;not null"`
	Team     Team   `json:"-"`
	PlayerID int    `json:"playerId" gorm:"index:idx_team_players_team_player_id,unique:not null"`
	Player   Player `json:"-"`
}
