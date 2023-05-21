package routes

import (
	"linkedin/controllers"
	"linkedin/middleware"

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
		api.GET("/profile", middleware.CheckAuth, controllers.Profile)
		api.GET("/users/:username", controllers.GetUser)
		api.GET("/posts", controllers.GetPosts)
		api.POST("/posts", middleware.CheckAuth, controllers.CreatePost)
	}
}
