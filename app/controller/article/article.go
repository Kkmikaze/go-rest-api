package article

import (
	"net/http"

	"github.com/Kkmikaze/go-rest-api/domain/article"
	"github.com/Kkmikaze/go-rest-api/domain/article/model"
	"github.com/Kkmikaze/go-rest-api/domain/article/repository"
	authLib "github.com/Kkmikaze/go-rest-api/lib/auth"
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

	authUser, err := authLib.GetAuthUserCtx(context)
	if err != nil {
		response.Error(context, http.StatusUnauthorized, err.Error())
		return
	}

	reqBody.UserID = authUser.ID

	if err := context.ShouldBind(&reqBody); err != nil {
		response.Error(context, http.StatusBadRequest, err.Error())
		return
	}
	if errStatus, err := ac.ArticleService.Create(&reqBody); err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, "Success create article", nil)
}

func (ac *articleController) Index(context *gin.Context) {
	resBody, errStatus, err := ac.ArticleService.Index()
	if err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, "Success get all articles", resBody)
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
	response.Json(context, http.StatusOK, "Success get article", resBody)
}

func (ac *articleController) Update(context *gin.Context) {
	slug := context.Param("slug")
	if slug == "" {
		response.Error(context, http.StatusBadRequest, "Slug is required")
		return
	}

	var reqBody model.ReqBodyUpdateArticle

	authUser, err := authLib.GetAuthUserCtx(context)
	if err != nil {
		response.Error(context, http.StatusUnauthorized, err.Error())
		return
	}

	reqBody.UserID = authUser.ID

	if err := context.ShouldBind(&reqBody); err != nil {
		response.Error(context, http.StatusBadRequest, err.Error())
		return
	}

	resBody, errStatus, err := ac.ArticleService.Update(slug, &reqBody)
	if err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}

	response.Json(context, http.StatusOK, "Success update article", resBody)
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

	response.Json(context, http.StatusOK, "Success delete article", nil)
}
