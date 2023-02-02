package user

import (
	"net/http"

	"github.com/Kkmikaze/go-rest-api/internal/domain/user/entity"
	"github.com/Kkmikaze/go-rest-api/internal/domain/user/repository"
	"github.com/Kkmikaze/go-rest-api/internal/domain/user/service"
	"github.com/Kkmikaze/go-rest-api/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authController struct {
	AuthService service.AuthServiceInterface
}

func AuthController(db *gorm.DB) *authController {
	return &authController{
		AuthService: service.AuthService(repository.AuthRepository(db)),
	}
}

func (ac *authController) Register(context *gin.Context) {
	var reqBody entity.ReqBodyRegister
	if err := context.ShouldBind(&reqBody); err != nil {
		response.Error(context, http.StatusBadRequest, err.Error())
		return
	}
	if errStatus, err := ac.AuthService.Register(&reqBody); err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, "Register success", nil)
}

func (ac *authController) Login(context *gin.Context) {
	var reqBody entity.ReqBodyLogin
	if err := context.ShouldBind(&reqBody); err != nil {
		response.Error(context, http.StatusBadRequest, err.Error())
		return
	}
	resBody, errStatus, err := ac.AuthService.Login(&reqBody)
	if err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, "Login success", resBody)
}

func (ac *authController) Logout(context *gin.Context) {
	user, err := ac.AuthService.CheckAuth(context.Request.Header.Get("Authorization"))
	if err != nil {
		response.Error(context, http.StatusUnauthorized, err.Error())
		return
	}
	if errStatus, err := ac.AuthService.Logout(*user); err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, "Logout success", nil)
}
