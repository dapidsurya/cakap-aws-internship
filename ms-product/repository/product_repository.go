package repository

import (
	"github.com/dapidsurya/cakap-aws-internship/ms-user/model/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindUserByUserId(userId int) []entity.Product
	CreateProduct(product *entity.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func InitProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindUserByUserId(userId int) []entity.Product {
	var products []entity.Product
	r.db.Where("user_id = ?", userId).Find(&products)
	return products
}

func (r *productRepository) CreateProduct(product *entity.Product) error {
	return r.db.Create(product).Error
}
