package teams

func FindByID(id string) (*Team, error) {
	stmt, err := findByIDStmt()
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)
	return scan(row)
}
