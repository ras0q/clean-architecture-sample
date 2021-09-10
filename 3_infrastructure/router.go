package infrastructure

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouting() {
	// API injection
	api, err := InjectAPIServer()
	if err != nil {
		log.Fatalf("failed to inject API: %v", err)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	echoAPI := e.Group("/api")

	pingAPI := echoAPI.Group("/ping")
	pingAPI.GET("", f(api.Ping.Ping))

	userAPI := echoAPI.Group("/users")
	userAPI.GET("", f(api.User.GetAll))
	userAPI.GET("/:id", f(api.User.GetByID))
	userAPI.POST("", f(api.User.Register))

	e.Logger.Fatal(e.Start(":8080"))
}
