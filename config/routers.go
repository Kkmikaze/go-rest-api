package config

import (
	"github.com/Kkmikaze/go-rest-api/app/controllers/root"
	"github.com/Kkmikaze/go-rest-api/config/collection"
	"github.com/Kkmikaze/go-rest-api/db"
	"github.com/gin-gonic/gin"
)

var Routers = gin.Default()

func init() {
	corsConfig(Routers)

	Routers.GET("/", root.Index)
	main := Routers.Group("v1")
	collection.MainRouter(db.DB, main)
}
