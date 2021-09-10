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

	userAPI := echoAPI.Group("/users")
	userAPI.GET("/", api.User.GetAll)
	userAPI.GET("/:id", api.User.GetByID)
	userAPI.POST("/", api.User.Register)
}
