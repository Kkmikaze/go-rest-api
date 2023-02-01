package middleware

import (
	"encoding/json"
	"net/http"

	auth "github.com/Kkmikaze/go-rest-api/domain/user"
	"github.com/Kkmikaze/go-rest-api/domain/user/repository"
	authLib "github.com/Kkmikaze/go-rest-api/lib/auth"
	"github.com/Kkmikaze/go-rest-api/lib/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authService := auth.AuthService(repository.AuthRepository(db))
		user, err := authService.CheckAuth(c.Request.Header.Get("Authorization"))
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		userStr, err := json.Marshal(&authLib.AuthData{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Set("auth", string(userStr))
	}
}
