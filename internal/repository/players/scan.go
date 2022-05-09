package players

type scanable interface {
	Scan(dest ...interface{}) error
}

func scan(s scanable) (*Player, error) {
	var player Player

	err := s.Scan(
		&player.ID, &player.Name, &player.Number,
		&player.CreatedAt, &player.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &player, nil
}
