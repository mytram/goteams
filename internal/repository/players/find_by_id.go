package players

func FindByID(id string) (*Player, error) {
	stmt, err := findByIDStmt()
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)
	return scan(row)
}
