package auth

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthData struct {
	ID        string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAuthUserCtx(ctx *gin.Context) (*AuthData, error) {
	authDataStr := ctx.GetString("auth")

	var authData AuthData
	err := json.Unmarshal([]byte(authDataStr), &authData)
	if err != nil {
		return nil, err
	}

	return &authData, nil
}
