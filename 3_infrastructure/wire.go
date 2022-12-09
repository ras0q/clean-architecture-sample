//go:build wireinject
// +build wireinject

package infrastructure

import (
	"github.com/google/wire"
	"github.com/ras0q/clean-architecture-sample/1_usecase/service"
	"github.com/ras0q/clean-architecture-sample/2_interface/handler"
	"github.com/ras0q/clean-architecture-sample/2_interface/repository"
)

// Note: エンドポイントを増やした時はここに追加する
var (
	apiSet  = wire.NewSet(handler.NewAPI)
	sqlSet  = wire.NewSet(NewSQLHandler)
	pingSet = wire.NewSet(handler.NewPingHandler)
	userSet = wire.NewSet(repository.NewUserRepository, service.NewUserService, handler.NewUserHandler)
)

func InjectAPIServer() (handler.API, error) {
	wire.Build(apiSet, sqlSet, pingSet, userSet)

	return handler.API{}, nil
}
