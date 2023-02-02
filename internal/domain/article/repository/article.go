package repository

import (
	"time"

	"github.com/Kkmikaze/go-rest-api/internal/domain/article/entity"
	"gorm.io/gorm"
)

type ArticleRepositoryInterface interface {
	FirstBySlug(slug *string) (entity.Article, error)
	IsAvailableSlug(slug string) bool
	All() ([]entity.Article, error)
	Create(reqBody *entity.ReqBodyCreateArticle) error
	Update(article entity.Article, reqBody *entity.ReqBodyUpdateArticle) error
	Delete(article entity.Article) error
}

type articleRepository struct {
	DB *gorm.DB
}

func ArticleRepository(DB *gorm.DB) ArticleRepositoryInterface {
	return &articleRepository{
		DB: DB,
	}
}

func (r *articleRepository) FirstBySlug(slug *string) (article entity.Article, err error) {
	if err := r.DB.Preload("User").Where("slug = ?", slug).First(&article).Error; err != nil {
		return article, err
	}
	return article, nil
}

func (r *articleRepository) IsAvailableSlug(slug string) bool {
	var article entity.Article
	if err := r.DB.Where("slug = ?", slug).Find(&article).RowsAffected; err != 0 {
		return false
	}
	return true
}

func (r *articleRepository) All() ([]entity.Article, error) {
	var articles []entity.Article
	if err := r.DB.Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

func (r *articleRepository) Create(reqBody *entity.ReqBodyCreateArticle) error {
	article := entity.Article{
		Title:  reqBody.Title,
		Body:   reqBody.Body,
		UserID: reqBody.UserID,
	}
	return r.DB.Create(&article).Error
}

func (r *articleRepository) Update(article entity.Article, reqBody *entity.ReqBodyUpdateArticle) error {
	return r.DB.Model(&article).Updates(&reqBody).Update("updated_at", time.Now()).Error
}

func (r *articleRepository) Delete(article entity.Article) error {
	return r.DB.Delete(&article).Error
}
