package service

import (
	"errors"
	"net/http"

	"github.com/Kkmikaze/go-rest-api/internal/domain/article/entity"
	"github.com/Kkmikaze/go-rest-api/internal/domain/article/repository"
	"github.com/Kkmikaze/go-rest-api/pkg/constant"
)

type ArticleServiceInterface interface {
	Create(reqBody *entity.ReqBodyCreateArticle) (int, error)
	Index() ([]entity.ResBodyArticle, int, error)
	Show(slug string) (*entity.ResBodyArticleDetail, int, error)
	Update(slug string, userID string, reqBody *entity.ReqBodyUpdateArticle) (*entity.ResBodyArticleDetail, int, error)
	Delete(slug string, userID string) (int, error)
}

type articleService struct {
	Repository repository.ArticleRepositoryInterface
}

func ArticleService(repository repository.ArticleRepositoryInterface) ArticleServiceInterface {
	return &articleService{
		Repository: repository,
	}
}

func (s *articleService) Create(reqBody *entity.ReqBodyCreateArticle) (int, error) {
	slug := entity.Slugify(reqBody.Title)
	checkArticle := s.Repository.IsAvailableSlug(slug)
	if !checkArticle {
		return http.StatusConflict, errors.New("article already exists")
	}
	if err := s.Repository.Create(reqBody); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func (s *articleService) Index() ([]entity.ResBodyArticle, int, error) {
	articles, err := s.Repository.All()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	var resBody []entity.ResBodyArticle
	for _, article := range articles {
		resBody = append(resBody, entity.ResBodyArticle{
			ID:        article.ID,
			Author:    article.User.Name,
			Title:     article.Title,
			Slug:      article.Slug,
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
		})
	}
	return resBody, http.StatusOK, nil
}

func (s *articleService) Show(slug string) (*entity.ResBodyArticleDetail, int, error) {
	article, err := s.Repository.FirstBySlug(&slug)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if article.ID == "" {
		return nil, http.StatusNotFound, errors.New(constant.RecordNotFound)
	}
	resBody := &entity.ResBodyArticleDetail{
		ID:        article.ID,
		Author:    article.User.Name,
		Title:     article.Title,
		Slug:      article.Slug,
		Body:      article.Body,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
	return resBody, http.StatusOK, nil
}

func (s *articleService) Update(slug string, userID string, reqBody *entity.ReqBodyUpdateArticle) (*entity.ResBodyArticleDetail, int, error) {
	article, err := s.Repository.FirstBySlug(&slug)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New(constant.RecordNotFound)
	}

	if userID != article.UserID {
		return nil, http.StatusUnauthorized, errors.New(constant.NotAuthorize)
	}

	if err := s.Repository.Update(article, reqBody); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	updateArticle, err := s.Repository.FirstBySlug(&article.Slug)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	resBody := &entity.ResBodyArticleDetail{
		ID:        updateArticle.ID,
		Author:    updateArticle.User.Name,
		Title:     updateArticle.Title,
		Slug:      updateArticle.Slug,
		Body:      updateArticle.Body,
		CreatedAt: updateArticle.CreatedAt,
		UpdatedAt: updateArticle.UpdatedAt,
	}
	return resBody, http.StatusOK, nil
}

func (s *articleService) Delete(slug string, userID string) (int, error) {
	article, err := s.Repository.FirstBySlug(&slug)
	if err != nil {
		return http.StatusBadRequest, errors.New(constant.RecordNotFound)
	}

	if userID != article.UserID {
		return http.StatusUnauthorized, errors.New(constant.NotAuthorize)
	}

	if err := s.Repository.Delete(article); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}
