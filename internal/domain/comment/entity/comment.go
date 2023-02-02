package entity

import (
	"time"

	"github.com/Kkmikaze/go-rest-api/internal/domain/user/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        string      `gorm:"type:char(36);primary_key"`
	ArticleID string      `gorm:"type:char(36);not null"`
	UserID    string      `gorm:"type:char(36);not null"`
	User      entity.User `gorm:"foreignKey:UserID"`
	Content   string      `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	comment.ID = uuid.New().String()
	return
}
