package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:3000", os.Getenv("CLIENT_URL")}
	router.Use(cors.New(config))
}

func StartRoute() {
	mapUrls()

	log.Info("Starting server")
	router.Run(":" + os.Getenv("PORT"))

}