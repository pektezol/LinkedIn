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
		api.POST("/posts/:post/comment", middleware.CheckAuth, controllers.CreateComment)
		api.DELETE("/comment/:id", middleware.CheckAuth, controllers.DeleteComment)
		api.GET("/connections", middleware.CheckAuth, controllers.GetConnectionRequests)
		api.POST("/connections/:username", middleware.CheckAuth, controllers.SendConnectionRequest)
		api.PUT("/connections/:username", middleware.CheckAuth, controllers.AcceptConnectionRequest)
		api.DELETE("/connections/:username", middleware.CheckAuth, controllers.RemoveConnection)
		api.POST("/like/:post", middleware.CheckAuth, controllers.Like)
		api.POST("/education", middleware.CheckAuth, controllers.AddEducation)
		api.DELETE("/education/:id", middleware.CheckAuth, controllers.DeleteEducation)
		api.POST("/experience", middleware.CheckAuth, controllers.AddExperience)
		api.DELETE("/experience/:id", middleware.CheckAuth, controllers.DeleteExperience)
		api.POST("/company", middleware.CheckAuth, controllers.CreateCompany)
	}
}
