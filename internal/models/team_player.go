package models

type TeamPlayer struct {
	Base
	TeamID   int `json:"teamId"`
	Team     Team
	PlayerID int `json:"playerId"`
	Player   Player
}
