package collection

import (
	"github.com/Kkmikaze/go-rest-api/internal/controller/article"
	"github.com/Kkmikaze/go-rest-api/internal/controller/comment"
	"github.com/Kkmikaze/go-rest-api/internal/controller/user"
	"github.com/Kkmikaze/go-rest-api/internal/middleware"
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

		commentCtrl := comment.CommentController(db)
		commentRoute := articleRoute.Group("/:slug/comment")
		{
			commentRoute.GET("", commentCtrl.AllByArticleSlug)
			commentRoute.POST("", middleware.Auth(db), commentCtrl.Create)
			commentRoute.PUT("/:id", middleware.Auth(db), commentCtrl.Update)
			commentRoute.DELETE("/:id", middleware.Auth(db), commentCtrl.Delete)
		}
	}
}
