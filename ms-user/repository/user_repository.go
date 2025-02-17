package repository

import (
	"errors"

	"github.com/dapidsurya/cakap-aws-internship/ms-user/config"
	"github.com/dapidsurya/cakap-aws-internship/ms-user/model/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{DB: config.GetDB()}
}

func (r *UserRepository) FindUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	result := r.DB.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, result.Error
}

func (r *UserRepository) FindAllUser() []entity.User {
	var users []entity.User
	r.DB.Find(&users)
	return users
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	return r.DB.Create(user).Error
}
