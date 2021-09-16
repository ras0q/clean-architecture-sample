package alt

import (
	"log"

	"github.com/gin-gonic/gin"
)

func InitRouting() *gin.Engine {
	// API injection
	api, err := InjectAPIServer()
	if err != nil {
		log.Fatalf("failed to inject API: %v", err)
	}

	r := gin.Default()

	ginAPI := r.Group("/api")

	pingAPI := ginAPI.Group("/ping")
	pingAPI.GET("", f(api.Ping.Ping))

	userAPI := ginAPI.Group("/users")
	userAPI.GET("", f(api.User.GetAll))
	userAPI.GET("/:id", f(api.User.GetByID))
	userAPI.POST("", f(api.User.Register))

	return r
}
