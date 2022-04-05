package teams

import (
	"encoding/json"
)

type scanable interface {
	Scan(dest ...interface{}) error
}

func scan(s scanable) (*Team, error) {
	var team Team
	var playersData []byte

	err := s.Scan(&team.ID, &team.Name, &team.CreatedAt, &team.UpdatedAt, &playersData)
	if err != nil {
		return nil, err
	}

	if len(playersData) > 0 {
		var players []*Player
		err = json.Unmarshal(playersData, &players)
		if err != nil {
			return nil, err
		}

		team.Players = players
	}

	return &team, nil
}
