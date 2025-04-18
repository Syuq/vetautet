// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/Syuq/go-vetautet-backend-api/internal/controller"
	"github.com/Syuq/go-vetautet-backend-api/internal/repo"
	"github.com/Syuq/go-vetautet-backend-api/internal/service"
)

// Injectors from user.wire.go:

func InitUserRouterHanlder() (*controller.UserController, error) {
	iUserRepository := repo.NewUserRepository()
	iUserAuthRepository := repo.NewUserAuthRepository()
	iUserService := service.NewUserService(iUserRepository, iUserAuthRepository)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}
