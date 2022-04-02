package handlers

import (
	"github.com/gin-gonic/gin"

	"goteams/internal/rest"
	"goteams/internal/teams/dtos"
	s "goteams/internal/teams/services"
)

func update(ctx *gin.Context) {
	id := ctx.Param("id")

	var dto dtos.UpdateDTO

	err := ctx.ShouldBind(&dto)
	if err != nil {
		rest.JSON(ctx, nil, rest.ErrBadQuest, err)
		return
	}

	team, err := s.Update(id, &dto)
	rest.JSON(ctx, team, err)
}
