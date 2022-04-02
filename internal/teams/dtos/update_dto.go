package dtos

type UpdateDTO struct {
	Name      string `json:"name"`
	PlayerIDs []uint `json:"playerIds"`
}
