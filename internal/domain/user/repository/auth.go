package repository

import (
	"github.com/Kkmikaze/go-rest-api/internal/domain/user/entity"
	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	FirstByEmail(email *string) (entity.User, error)
	FirstByID(string) (entity.User, error)
	Create(reqBody *entity.ReqBodyRegister) error
	UpdateToken(entity.User, string) error
	DeleteToken(entity.User) error
}

type authRepository struct {
	DB *gorm.DB
}

func AuthRepository(DB *gorm.DB) AuthRepositoryInterface {
	return &authRepository{
		DB: DB,
	}
}

func (r *authRepository) FirstByEmail(email *string) (user entity.User, err error) {
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *authRepository) FirstByID(id string) (entity.User, error) {
	var user entity.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *authRepository) Create(reqBody *entity.ReqBodyRegister) error {
	user := entity.User{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: reqBody.Password,
	}
	return r.DB.Create(&user).Error
}

func (r *authRepository) UpdateToken(user entity.User, ss string) error {
	user.Token = ss
	return r.DB.Save(&user).Error
}

func (r *authRepository) DeleteToken(user entity.User) error {
	user.Token = ""
	return r.DB.Save(&user).Error
}
