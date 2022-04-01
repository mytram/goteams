package dtos

type CreateDTO struct {
	Name   string `json:"name" binding:"required"`
	Number uint   `json:"number" binding:"required"`
}
