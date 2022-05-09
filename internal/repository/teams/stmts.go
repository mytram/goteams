package teams

import (
	"database/sql"

	"goteams/internal/db"
)

var stmtFind *sql.Stmt
var stmtFindByID *sql.Stmt
var stmtCreate *sql.Stmt

func createStmt() (*sql.Stmt, error) {
	if stmtCreate != nil {
		return stmtCreate, nil
	}

	dbh, err := db.New()
	if err != nil {
		return nil, err
	}

	stmtCreate, err := dbh.Prepare(sqlInsertTeam)
	return stmtCreate, err
}

func findStmt() (*sql.Stmt, error) {
	if stmtFind != nil {
		return stmtFind, nil
	}

	dbh, err := db.New()
	if err != nil {
		return nil, err
	}

	stmtFind, err := dbh.Prepare(sqlSelectTeams)
	return stmtFind, err
}

func findByIDStmt() (*sql.Stmt, error) {
	if stmtFindByID != nil {
		return stmtFindByID, nil
	}

	dbh, err := db.New()
	if err != nil {
		return nil, err
	}

	stmtFindByID, err := dbh.Prepare(sqlSelectTeamById)
	return stmtFindByID, err
}
