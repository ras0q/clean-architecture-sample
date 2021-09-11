// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package infrastructure

import (
	"github.com/Ras96/clean-architecture-sample/1_usecase"
	"github.com/Ras96/clean-architecture-sample/2_interface/database"
	"github.com/Ras96/clean-architecture-sample/2_interface/handler"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InjectAPIServer() (handler.API, error) {
	pingHandler := handler.NewPingHandler()
	databaseSQLHandler := NewSQLHandler()
	userRepository := database.NewUserRepository(databaseSQLHandler)
	userService := usecase.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	api := handler.NewAPI(pingHandler, userHandler)
	return api, nil
}

// wire.go:

// エンドポイントを増やした時はここに追加する
var (
	apiSet  = wire.NewSet(handler.NewAPI)
	sqlSet  = wire.NewSet(NewSQLHandler)
	pingSet = wire.NewSet(handler.NewPingHandler)
	userSet = wire.NewSet(database.NewUserRepository, usecase.NewUserService, handler.NewUserHandler)
)
