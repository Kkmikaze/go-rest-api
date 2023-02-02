package collection

import (
	"github.com/Kkmikaze/go-rest-api/app/controller/article"
	"github.com/Kkmikaze/go-rest-api/app/controller/user"
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

	articleCtrl := article.ArticleController(db)
	articleRoute := main.Group("article")
	{
		articleRoute.POST("", middleware.Auth(db), articleCtrl.Create)
		articleRoute.GET("", articleCtrl.Index)
		articleRoute.GET("/:slug", articleCtrl.Show)
		articleRoute.PUT("/:slug", middleware.Auth(db), articleCtrl.Update)
		articleRoute.DELETE("/:slug", middleware.Auth(db), articleCtrl.Delete)
	}
}
