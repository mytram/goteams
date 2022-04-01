package handlers

import (
	"github.com/gin-gonic/gin"

	"goteams/internal/players/dtos"
	s "goteams/internal/players/services"
	"goteams/internal/rest"
)

func create(ctx *gin.Context) {
	var dto dtos.CreateDTO

	err := ctx.ShouldBind(&dto)
	if err != nil {
		rest.JSON(ctx, nil, rest.ErrBadQuest, err)
		return
	}

	player, err := s.Create(&dto)
	rest.JSON(ctx, player, err)
}
