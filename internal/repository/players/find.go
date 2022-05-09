package players

func Find() ([]Player, error) {
	stmt, err := findStmt()
	if err != nil {
		return nil, err
	}

	var players []Player

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		player, err := scan(rows)
		if err != nil {
			return nil, err
		}

		players = append(players, *player)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return players, nil
}
