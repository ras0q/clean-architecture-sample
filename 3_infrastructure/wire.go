//go:build wireinject
// +build wireinject

package infrastructure

import (
	"github.com/Ras96/clean-architecture-sample/1_usecase/service"
	"github.com/Ras96/clean-architecture-sample/2_interface/handler"
	"github.com/Ras96/clean-architecture-sample/2_interface/repository"
	"github.com/google/wire"
)

// エンドポイントを増やした時はここに追加する
var (
	apiSet  = wire.NewSet(handler.NewAPI)
	sqlSet  = wire.NewSet(NewSQLHandler)
	userSet = wire.NewSet(repository.NewUserRepository, service.NewUserService, handler.NewUserHandler)
)

func InjectAPIServer() (handler.API, error) {
	wire.Build(apiSet, sqlSet, userSet)

	return handler.API{}, nil
}
