package article

import (
	"net/http"

	"github.com/Kkmikaze/go-rest-api/internal/domain/article/entity"
	"github.com/Kkmikaze/go-rest-api/internal/domain/article/repository"
	"github.com/Kkmikaze/go-rest-api/internal/domain/article/service"
	authLib "github.com/Kkmikaze/go-rest-api/pkg/auth"
	"github.com/Kkmikaze/go-rest-api/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type articleController struct {
	ArticleService service.ArticleServiceInterface
}

func ArticleController(db *gorm.DB) *articleController {
	return &articleController{
		ArticleService: service.ArticleService(repository.ArticleRepository(db)),
	}
}

func (ac *articleController) Create(context *gin.Context) {
	var reqBody entity.ReqBodyCreateArticle

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

	var reqBody entity.ReqBodyUpdateArticle
	if err := context.ShouldBind(&reqBody); err != nil {
		response.Error(context, http.StatusBadRequest, err.Error())
		return
	}

	authUser, err := authLib.GetAuthUserCtx(context)
	if err != nil {
		response.Error(context, http.StatusUnauthorized, err.Error())
		return
	}

	resBody, errStatus, err := ac.ArticleService.Update(slug, authUser.ID, &reqBody)
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

	authUser, err := authLib.GetAuthUserCtx(context)
	if err != nil {
		response.Error(context, http.StatusUnauthorized, err.Error())
		return
	}

	if errStatus, err := ac.ArticleService.Delete(slug, authUser.ID); err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}

	response.Json(context, http.StatusOK, "Success delete article", nil)
}
