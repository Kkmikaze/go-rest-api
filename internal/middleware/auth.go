package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/Kkmikaze/go-rest-api/internal/domain/user/repository"
	"github.com/Kkmikaze/go-rest-api/internal/domain/user/service"
	"github.com/Kkmikaze/go-rest-api/pkg/auth"
	"github.com/Kkmikaze/go-rest-api/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authService := service.AuthService(repository.AuthRepository(db))
		user, err := authService.CheckAuth(c.Request.Header.Get("Authorization"))
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		userStr, err := json.Marshal(&auth.AuthData{
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
