package handlers

import (
	"github.com/gin-gonic/gin"

	"goteams/internal/rest"
	s "goteams/internal/teams/services"
)

func getTeams(ctx *gin.Context) {
	teams, err := s.Fetch()
	rest.JSON(ctx, teams, err)
}

func getTeam(ctx *gin.Context) {
	id := ctx.Param("id")

	team, err := s.FetchById(id)

	rest.JSON(ctx, team, err)
}
