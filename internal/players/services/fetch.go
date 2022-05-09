package services

import (
	repo "goteams/internal/repository/players"
)

func Fetch() ([]Player, error) {
	return repo.Find()
}

func FetchById(id string) (*Player, error) {
	return repo.FindByID(id)
}
