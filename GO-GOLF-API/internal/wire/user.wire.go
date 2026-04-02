//go:build wireinject
// +build wireinject

package wire

import (
	"GO-GOLF-API/internal/controller"
	"GO-GOLF-API/internal/repo"
	service "GO-GOLF-API/internal/service"
	service_impl "GO-GOLF-API/internal/service/impl"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service_impl.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}

func InitUserServiceHandler() (service.IUserService, error) {
	wire.Build(
		repo.NewUserRepository,
		service_impl.NewUserService,
	)

	return nil, nil
}
