package comment

import (
	"net/http"

	ar "github.com/Kkmikaze/go-rest-api/internal/domain/article/repository"
	as "github.com/Kkmikaze/go-rest-api/internal/domain/article/service"
	"github.com/Kkmikaze/go-rest-api/internal/domain/comment/entity"
	cr "github.com/Kkmikaze/go-rest-api/internal/domain/comment/repository"
	cs "github.com/Kkmikaze/go-rest-api/internal/domain/comment/service"
	authLib "github.com/Kkmikaze/go-rest-api/pkg/auth"
	"github.com/Kkmikaze/go-rest-api/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentController struct {
	CommentService cs.CommentServiceInterface
	ArticleService as.ArticleServiceInterface
}

func CommentController(db *gorm.DB) *commentController {
	return &commentController{
		CommentService: cs.CommentService(cr.CommentRepository(db)),
		ArticleService: as.ArticleService(ar.ArticleRepository(db)),
	}
}

func (c *commentController) AllByArticleSlug(context *gin.Context) {
	resArticle, errStatus, err := c.ArticleService.Show(context.Param("slug"))
	if err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}

	resComments, err := c.CommentService.Index(resArticle.ID)
	if err != nil {
		response.Error(context, http.StatusInternalServerError, err.Error())
		return
	}
	response.Json(context, http.StatusOK, "Success get comments", resComments)
}

func (c *commentController) Create(context *gin.Context) {
	authUser, err := authLib.GetAuthUserCtx(context)
	if err != nil {
		response.Error(context, http.StatusUnauthorized, err.Error())
		return
	}

	reqBody := entity.ReqBodyCreateComment{
		UserID:      string(authUser.ID),
		ArticleSlug: context.Param("slug"),
	}
	if err := context.ShouldBind(&reqBody); err != nil {
		response.Error(context, http.StatusBadRequest, err.Error())
		return
	}

	resArticle, errStatus, err := c.ArticleService.Show(reqBody.ArticleSlug)

	if err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}

	if errStatus, err := c.CommentService.Create(resArticle.ID, &reqBody); err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, "Success create comment", nil)
}

func (c *commentController) Update(context *gin.Context) {
	authUser, err := authLib.GetAuthUserCtx(context)
	if err != nil {
		response.Error(context, http.StatusUnauthorized, err.Error())
		return
	}

	reqBody := entity.ReqBodyUpdateComment{
		ID:          context.Param("id"),
		ArticleSlug: context.Param("slug"),
		UserID:      string(authUser.ID),
	}
	if err := context.ShouldBind(&reqBody); err != nil {
		response.Error(context, http.StatusBadRequest, err.Error())
		return
	}

	resArticle, errStatus, err := c.ArticleService.Show(reqBody.ArticleSlug)

	if err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}

	if errStatus, err := c.CommentService.Update(resArticle.ID, &reqBody); err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}

	response.Json(context, http.StatusOK, "Success update comment", nil)
}

func (c *commentController) Delete(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		response.Error(context, http.StatusBadRequest, "Comment ID is required")
		return
	}

	authUser, err := authLib.GetAuthUserCtx(context)
	if err != nil {
		response.Error(context, http.StatusUnauthorized, err.Error())
		return
	}

	if errStatus, err := c.CommentService.Delete(id, authUser.ID); err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}

	response.Json(context, http.StatusOK, "Success delete comment", nil)
}
