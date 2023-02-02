package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID        string
	Title     string
	Slug      string
	Body      string
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
