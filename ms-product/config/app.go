package config

import (
	"github.com/dapidsurya/cakap-aws-internship/ms-user/handler"
	"github.com/dapidsurya/cakap-aws-internship/ms-user/repository"
)

type App struct {
	ProductHandler handler.ProductHandler
}

func InitializeApp() *App {
	db := GetDB()
	productRepo := repository.InitProductRepository(db)
	productHandler := handler.InitProductHandler(productRepo)

	return &App{
		ProductHandler: productHandler,
	}
}
