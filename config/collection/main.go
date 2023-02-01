package collection

import (
	"github.com/Kkmikaze/go-rest-api/app/controllers/user"
	"github.com/Kkmikaze/go-rest-api/config/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MainRouter(db *gorm.DB, main *gin.RouterGroup) {
	userAuthCtrl := user.AuthController(db)
	authRoute := main.Group("auth")
	{
		authRoute.POST("/register", userAuthCtrl.Register)
		authRoute.POST("/login", userAuthCtrl.Login)
		authRoute.POST("/logout", middleware.Auth(db), userAuthCtrl.Logout)
	}
}
