package user

import (
	"errors"
	"net/http"
	"time"

	"github.com/Kkmikaze/go-rest-api/domain/user/model"
	"github.com/Kkmikaze/go-rest-api/domain/user/repository"
	"github.com/Kkmikaze/go-rest-api/lib/constant"
	"github.com/Kkmikaze/go-rest-api/lib/encrypt"
	"github.com/dgrijalva/jwt-go"
)

type AuthServiceInterface interface {
	Register(reqBody *model.ReqBodyRegister) (int, error)
	Login(reqBody *model.ReqBodyLogin) (*model.ResBody, int, error)
	CheckAuth(string) (*model.User, error)
	Logout(model.User) (int, error)
}

type authService struct {
	Repository repository.AuthRepositoryInterface
}

func AuthService(repository repository.AuthRepositoryInterface) AuthServiceInterface {
	return &authService{
		Repository: repository,
	}
}

func (s *authService) Register(reqBody *model.ReqBodyRegister) (int, error) {
	if _, err := s.Repository.FirstByEmail(&reqBody.Email); err == nil {
		return http.StatusBadRequest, errors.New(constant.EmailAlreadyExists)
	}
	if err := encrypt.GenerateFromPassword(&reqBody.Password); err != nil {
		return http.StatusInternalServerError, err
	}

	err := s.Repository.Create(reqBody)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *authService) Login(reqBody *model.ReqBodyLogin) (*model.ResBody, int, error) {
	user, err := s.Repository.FirstByEmail(&reqBody.Email)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New(constant.UserNotFound)
	}
	if err = encrypt.CompareHashAndPassword(&user.Password, &reqBody.Password); err != nil {
		return nil, http.StatusBadRequest, errors.New(constant.PasswordIsIncorrect)
	}
	claims := model.Jwt{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	ss, err := encrypt.NewWithClaims(&claims)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if err = s.Repository.UpdateToken(user, ss); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	var resBody model.ResBody
	resBody.Token = ss
	return &resBody, http.StatusOK, nil
}

func (s *authService) CheckAuth(bearerToken string) (*model.User, error) {
	tokenRaw, claims, err := encrypt.Parse(bearerToken)
	if err != nil {
		return nil, err
	}
	id := string(claims["id"].(string))
	user, err := s.Repository.FirstByID(id)
	if err != nil {
		if err.Error() == constant.RecordNotFound {
			return nil, errors.New(constant.UserNotFound)
		}
	}
	if user.Token != tokenRaw {
		return nil, errors.New(constant.UserHasSignedOut)
	}
	return &user, nil
}

func (s *authService) Logout(user model.User) (int, error) {
	if err := s.Repository.DeleteToken(user); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
