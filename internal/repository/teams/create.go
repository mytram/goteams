package teams

func Create(team *Team) (*Team, error) {
	stmt, err := createStmt()
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(team.Name)
	return scan(row)
}
