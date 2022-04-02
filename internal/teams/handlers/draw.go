package handlers

import (
	"github.com/gin-gonic/gin"
)

func Draw(r *gin.Engine) {
	r.GET("/teams", getTeams)
	r.GET("/teams/:id", getTeam)
	r.POST("/teams", create)
	r.PUT("/teams/:id", update)
}
