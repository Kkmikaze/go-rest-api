package config

import (
	"github.com/Kkmikaze/go-rest-api/config/db"
	"github.com/Kkmikaze/go-rest-api/internal/controller/root"
	router "github.com/Kkmikaze/go-rest-api/internal/router"
	"github.com/gin-gonic/gin"
)

var Routers = gin.Default()

func init() {
	corsConfig(Routers)

	Routers.GET("/", root.Index)
	main := Routers.Group("v1")
	router.MainRouter(db.DB, main)
}
