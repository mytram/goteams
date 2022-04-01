package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	players "goteams/internal/players/handlers"
	"goteams/internal/repository"
)

func main() {
	os.Setenv("TZ", "UTC")
	repository.Setup(os.Getenv("DB_FILE"))

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
	}))

	players.Draw(r)

	r.Run()
}