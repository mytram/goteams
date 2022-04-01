package models

type Player struct {
	Base
	Name   string `json:"name"`
	Number uint   `json:"number" gorm:"unique"`
}
