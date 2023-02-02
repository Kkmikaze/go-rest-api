package repository

import (
	"time"

	"github.com/Kkmikaze/go-rest-api/domain/article/model"
	"gorm.io/gorm"
)

type ArticleRepositoryInterface interface {
	FirstBySlug(slug *string) (model.Article, error)
	IsAvailableSlug(slug string) bool
	All() ([]model.Article, error)
	Create(reqBody *model.ReqBodyCreateArticle) error
	Update(article model.Article, reqBody *model.ReqBodyUpdateArticle) error
	Delete(article model.Article) error
}

type articleRepository struct {
	DB *gorm.DB
}

func ArticleRepository(DB *gorm.DB) ArticleRepositoryInterface {
	return &articleRepository{
		DB: DB,
	}
}

func (r *articleRepository) FirstBySlug(slug *string) (article model.Article, err error) {
	if err := r.DB.Where("slug = ?", slug).First(&article).Error; err != nil {
		return article, err
	}
	return article, nil
}

func (r *articleRepository) IsAvailableSlug(slug string) bool {
	var article model.Article
	if err := r.DB.Where("slug = ?", slug).Find(&article).RowsAffected; err != 0 {
		return false
	}
	return true
}

func (r *articleRepository) All() ([]model.Article, error) {
	var articles []model.Article
	if err := r.DB.Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}

func (r *articleRepository) Create(reqBody *model.ReqBodyCreateArticle) error {
	var article model.Article
	article.Title = reqBody.Title
	article.Body = reqBody.Body
	return r.DB.Create(&article).Error
}

func (r *articleRepository) Update(article model.Article, reqBody *model.ReqBodyUpdateArticle) error {
	article.Title = reqBody.Title
	article.Body = reqBody.Body
	article.UpdatedAt = time.Now()
	return r.DB.Save(&article).Error
}

func (r *articleRepository) Delete(article model.Article) error {
	return r.DB.Delete(&article).Error
}
