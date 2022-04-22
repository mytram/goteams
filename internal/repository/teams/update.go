package teams

import (
	"context"
	"database/sql"
	"strings"

	"goteams/internal/db"
)

func Update(id string, name string, playerIDs *[]uint) (*Team, error) {
	dbh, err := db.New()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	tx, err := dbh.BeginTx(ctx, nil)
	defer tx.Rollback()

	err = updateTx(ctx, tx, id, name, playerIDs)
	if err != nil {

		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return FindByID(id)
}

func updateTx(ctx context.Context, tx *sql.Tx, id string, name string, playerIDs *[]uint) (err error) {
	if name := strings.TrimSpace(name); name != "" {
		_, err = tx.ExecContext(ctx, sqlUpdateTeamPlayer, id, name)
		if err != nil {
			return
		}
	}

	if playerIDs == nil {
		return
	}

	_, err = tx.ExecContext(ctx, sqlDeleteTeamPlayers, id)
	if err != nil {
		return
	}

	for _, playerID := range *playerIDs {
		_, err = tx.ExecContext(ctx, sqlInsertTeamPlayer, id, playerID)
		if err != nil {
			return
		}
	}

	return
}
