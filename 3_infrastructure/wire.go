//go:build wireinject
// +build wireinject

package infrastructure

import (
	usecase "github.com/Ras96/clean-architecture-sample/1_usecase"
	"github.com/Ras96/clean-architecture-sample/2_interface/database"
	"github.com/Ras96/clean-architecture-sample/2_interface/handler"
	"github.com/google/wire"
)

// エンドポイントを増やした時はここに追加する
var (
	apiSet  = wire.NewSet(handler.NewAPI)
	sqlSet  = wire.NewSet(NewSQLHandler)
	pingSet = wire.NewSet(handler.NewPingHandler)
	userSet = wire.NewSet(database.NewUserRepository, usecase.NewUserService, handler.NewUserHandler)
)

func InjectAPIServer() (handler.API, error) {
	wire.Build(apiSet, sqlSet, pingSet, userSet)

	return handler.API{}, nil
}
