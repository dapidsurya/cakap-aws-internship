package config

import (
	"github.com/dapidsurya/cakap-aws-internship/ms-user/handler"
	"github.com/dapidsurya/cakap-aws-internship/ms-user/repository"
)

type App struct {
	UserHandler handler.UserHandler
}

func InitializeApp() *App {
	db := GetDB()
	userRepo := repository.NewUserRepository(db)
	userHandler := handler.InitUserHandler(userRepo)

	return &App{
		UserHandler: userHandler,
	}
}
