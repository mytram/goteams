package db

import (
	"database/sql"
)

var dbh *sql.DB

func New() (*sql.DB, error) {
	if dbh != nil {
		return dbh, nil
	}

	var err error

	dbh, err = sql.Open("pgx", `postgresql://postgres:postgres@localhost:5432/goteams`)
	if err != nil {
		return nil, err
	}

	return dbh, nil
}
