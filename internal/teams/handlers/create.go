package handlers

import (
	"github.com/gin-gonic/gin"

	"goteams/internal/rest"
	"goteams/internal/teams/dtos"
	s "goteams/internal/teams/services"
)

func create(ctx *gin.Context) {
	var dto dtos.CreateDTO

	err := ctx.ShouldBind(&dto)
	if err != nil {
		rest.JSON(ctx, nil, rest.ErrBadQuest, err)
		return
	}

	team, err := s.Create(&dto)
	rest.JSON(ctx, team, err)
}
