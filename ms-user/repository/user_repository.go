package repository

import (
	"errors"

	"github.com/dapidsurya/cakap-aws-internship/ms-user/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUser() []entity.User
	FindUserByUsername(username string) (*entity.User, error)
	CreateUser(user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, result.Error
}

func (r *userRepository) FindAllUser() []entity.User {
	var users []entity.User
	r.db.Find(&users)
	return users
}

func (r *userRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}
