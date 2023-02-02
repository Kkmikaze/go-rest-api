package repository

import (
	"github.com/Kkmikaze/go-rest-api/internal/domain/comment/entity"
	"gorm.io/gorm"
)

type CommentRepositoryInterface interface {
	AllByArticleID(articleID string) ([]entity.Comment, error)
	FirstByID(id *string) (entity.Comment, error)
	Create(articleID string, reqBody *entity.ReqBodyCreateComment) error
	Update(comment entity.Comment, reqBody *entity.ReqBodyUpdateComment) error
	Delete(comment entity.Comment) error
}

type commentRepository struct {
	DB *gorm.DB
}

func CommentRepository(DB *gorm.DB) CommentRepositoryInterface {
	return &commentRepository{
		DB: DB,
	}
}

func (r *commentRepository) AllByArticleID(articleID string) ([]entity.Comment, error) {
	var comments []entity.Comment
	if err := r.DB.Preload("User").Where("article_id = ?", articleID).Find(&comments).Error; err != nil {
		return comments, err
	}
	return comments, nil
}

func (r *commentRepository) FirstByID(id *string) (comment entity.Comment, err error) {
	if err := r.DB.Preload("User").Where("id = ?", id).First(&comment).Error; err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *commentRepository) Create(articleID string, reqBody *entity.ReqBodyCreateComment) error {
	comment := entity.Comment{
		ArticleID: articleID,
		UserID:    reqBody.UserID,
		Content:   reqBody.Content,
	}
	return r.DB.Create(&comment).Error
}

func (r *commentRepository) Update(comment entity.Comment, reqBody *entity.ReqBodyUpdateComment) error {
	return r.DB.Model(&comment).Updates(entity.Comment{Content: reqBody.Content}).Error
}

func (r *commentRepository) Delete(comment entity.Comment) error {
	return r.DB.Delete(&comment).Error
}
