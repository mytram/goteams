package models

type Team struct {
	Base
	Name    string    `json:"name" gorm:"unique;"`
	Players []*Player `json:"players" gorm:"many2many:team_players;"`
}
