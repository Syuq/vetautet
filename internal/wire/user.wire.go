//go:build wireinject

package wire

import (
	"github.com/Syuq/go-vetautet-backend-api/internal/controller"
	"github.com/Syuq/go-vetautet-backend-api/internal/repo"
	"github.com/Syuq/go-vetautet-backend-api/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHanlder() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
