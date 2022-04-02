package models

type Player struct {
	Base
	Name   string `json:"name" gorm:"not null"`
	Number uint   `json:"number" gorm:"unique;not null"`
}
