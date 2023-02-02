package article

import (
	"net/http"

	"github.com/Kkmikaze/go-rest-api/domain/article"
	"github.com/Kkmikaze/go-rest-api/domain/article/model"
	"github.com/Kkmikaze/go-rest-api/domain/article/repository"
	"github.com/Kkmikaze/go-rest-api/lib/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type articleController struct {
	ArticleService article.ArticleServiceInterface
}

func ArticleController(db *gorm.DB) *articleController {
	return &articleController{
		ArticleService: article.ArticleService(repository.ArticleRepository(db)),
	}
}

func (ac *articleController) Create(context *gin.Context) {
	var reqBody model.ReqBodyCreateArticle
	if err := context.ShouldBind(&reqBody); err != nil {
		response.Error(context, http.StatusBadRequest, err.Error())
		return
	}
	if errStatus, err := ac.ArticleService.Create(&reqBody); err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, "Create success", nil)
}

func (ac *articleController) Index(context *gin.Context) {
	resBody, errStatus, err := ac.ArticleService.Index()
	if err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, "Index success", resBody)
}

func (ac *articleController) Show(context *gin.Context) {
	slug := context.Param("slug")
	if slug == "" {
		response.Error(context, http.StatusBadRequest, "Slug is required")
		return
	}

	resBody, errStatus, err := ac.ArticleService.Show(slug)
	if err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, "Show success", resBody)
}

func (ac *articleController) Update(context *gin.Context) {
	var reqBody model.ReqBodyUpdateArticle
	if err := context.ShouldBind(&reqBody); err != nil {
		response.Error(context, http.StatusBadRequest, err.Error())
		return
	}

	slug := context.Param("slug")
	if slug == "" {
		response.Error(context, http.StatusBadRequest, "Slug is required")
		return
	}

	resBody, errStatus, err := ac.ArticleService.Update(slug, &reqBody)
	if err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}

	response.Json(context, http.StatusOK, "Update success", resBody)
}

func (ac *articleController) Delete(context *gin.Context) {
	slug := context.Param("slug")
	if slug == "" {
		response.Error(context, http.StatusBadRequest, "Slug is required")
		return
	}

	if errStatus, err := ac.ArticleService.Delete(slug); err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}

	response.Json(context, http.StatusOK, "Delete success", nil)
}
