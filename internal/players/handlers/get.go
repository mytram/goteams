package handlers

import (
	"github.com/gin-gonic/gin"

	s "goteams/internal/players/services"
	"goteams/internal/rest"
)

func getPlayers(ctx *gin.Context) {
	players, err := s.Fetch()
	rest.JSON(ctx, players, err)
}

func getPlayer(ctx *gin.Context) {
	id := ctx.Param("id")

	player, err := s.FetchById(id)

	rest.JSON(ctx, player, err)
}
