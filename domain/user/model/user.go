package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return
}

type Jwt struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	*jwt.StandardClaims
}
