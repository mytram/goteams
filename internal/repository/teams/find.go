package teams

func Find() ([]Team, error) {
	stmt, err := findStmt()
	if err != nil {
		return nil, err
	}

	var teams []Team

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		team, err := scan(rows)
		if err != nil {
			return nil, err
		}

		teams = append(teams, *team)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return teams, nil
}
