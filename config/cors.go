package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsConfig(routers *gin.Engine) {
	routers.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Authorization", "Content-Type"},
		AllowMethods:    []string{"PUT", "PATCH", "DELETE", "POST", "OPTIONS", "GET"},
	}))
}
