package model

import (
	"strings"
	"time"

	"github.com/Kkmikaze/go-rest-api/domain/user/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID        string     `gorm:"type:char(36);primary_key"`
	UserID    string     `gorm:"type:char(36);not null"`
	User      model.User `gorm:"foreignKey:UserID"`
	Title     string     `gorm:"type:varchar(255);not null"`
	Slug      string     `gorm:"type:varchar(255);not null;unique"`
	Body      string     `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (article *Article) BeforeCreate(tx *gorm.DB) (err error) {
	article.ID = uuid.New().String()
	article.Slug = Slugify(article.Title)
	return
}

func Slugify(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}
