package players

func Create(player *Player) (*Player, error) {
	stmt, err := createStmt()
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(player.Name, player.Number)
	return scan(row)
}
