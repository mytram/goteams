package services

import (
	repo "goteams/internal/repository/teams"
)

func Fetch() (teams []Team, err error) {
	return repo.Find()
}

func FetchById(id string) (team *Team, err error) {
	return repo.FindByID(id)
}
