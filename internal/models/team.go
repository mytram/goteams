package models

type Team struct {
	Base
	Name string `json:"name" gorm:"unique"`
}
