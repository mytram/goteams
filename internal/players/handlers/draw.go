package handlers

import (
	"github.com/gin-gonic/gin"
)

func Draw(r *gin.Engine) {
	r.GET("/players", getPlayers)
	r.GET("/players/:id", getPlayer)
	r.POST("/players", create)
}
