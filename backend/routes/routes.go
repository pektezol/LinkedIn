package routes

import (
	"linkedin/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}
		api.GET("/profile", controllers.Profile)
		api.GET("/users/:username", controllers.GetUser)
	}
}
