package service

import (
	"errors"
	"net/http"

	"github.com/Kkmikaze/go-rest-api/internal/domain/comment/entity"
	"github.com/Kkmikaze/go-rest-api/internal/domain/comment/repository"
	"github.com/Kkmikaze/go-rest-api/pkg/constant"
)

type CommentServiceInterface interface {
	Index(articleID string) ([]entity.ResBodyComment, error)
	FirstByID(id *string) (entity.Comment, error)
	Create(articleID string, reqBody *entity.ReqBodyCreateComment) (int, error)
	Update(articleID string, reqBody *entity.ReqBodyUpdateComment) (int, error)
	Delete(id string, userID string) (int, error)
}

type commentService struct {
	Repository repository.CommentRepositoryInterface
}

func CommentService(repository repository.CommentRepositoryInterface) CommentServiceInterface {
	return &commentService{
		Repository: repository,
	}
}

func (s *commentService) Index(articleID string) ([]entity.ResBodyComment, error) {
	comments, err := s.Repository.AllByArticleID(articleID)
	if err != nil {
		return nil, err
	}
	var resBodyComments []entity.ResBodyComment
	for _, comment := range comments {
		resBodyComments = append(resBodyComments, entity.ResBodyComment{
			ID:        comment.ID,
			ArticleID: comment.ArticleID,
			WriteBy:   comment.User.Name,
			Content:   comment.Content,
			UpdateAt:  comment.UpdatedAt,
		})
	}
	return resBodyComments, nil
}

func (s *commentService) FirstByID(id *string) (comment entity.Comment, err error) {
	comment, err = s.Repository.FirstByID(id)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (s *commentService) Create(articleID string, reqBody *entity.ReqBodyCreateComment) (int, error) {
	if err := s.Repository.Create(articleID, reqBody); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func (s *commentService) Update(articleID string, reqBody *entity.ReqBodyUpdateComment) (int, error) {
	comment, err := s.Repository.FirstByID(&reqBody.ID)
	if err != nil {
		return http.StatusNotFound, err
	}

	if comment.UserID != reqBody.UserID {
		return http.StatusUnauthorized, errors.New(constant.NotAuthorize)
	}

	if comment.ArticleID != articleID {
		return http.StatusNotFound, errors.New(constant.RecordNotFound)
	}

	if err := s.Repository.Update(comment, reqBody); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *commentService) Delete(id string, userID string) (int, error) {
	comment, err := s.Repository.FirstByID(&id)
	if err != nil {
		return http.StatusNotFound, err
	}

	if comment.UserID != userID {
		return http.StatusUnauthorized, errors.New(constant.NotAuthorize)
	}

	if err := s.Repository.Delete(comment); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
