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
		api.PUT("/profile", middleware.CheckAuth, controllers.EditProfile)
		api.GET("/users/:username", controllers.GetUser)
		api.GET("/posts", middleware.CheckAuth, controllers.GetPosts)
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
		api.POST("/skill", middleware.CheckAuth, controllers.AddSkill)
		api.DELETE("/skill/:id", middleware.CheckAuth, controllers.DeleteSkill)
		api.POST("/cv", middleware.CheckAuth, controllers.AddCV)
		api.DELETE("/cv", middleware.CheckAuth, controllers.DeleteCV)
		api.GET("/company", controllers.GetCompany)
		api.POST("/company", middleware.CheckAuth, controllers.CreateCompany)
		api.GET("/company/:id/job", middleware.CheckAuth, controllers.GetJobApplications)
		api.POST("/company/:id/job", middleware.CheckAuth, controllers.CreateJobOpening)
		api.GET("/search", controllers.Search)
		api.GET("/messages", middleware.CheckAuth, controllers.GetAllMessages)
		api.GET("/messages/:username", middleware.CheckAuth, controllers.GetSpecificMessage)
		api.POST("/messages/:username", middleware.CheckAuth, controllers.SendMessage)
		api.GET("/notifications", middleware.CheckAuth, controllers.GetNotifications)
		api.GET("/jobs", middleware.CheckAuth, controllers.GetJobOpenings)
		api.POST("/jobs/:id", middleware.CheckAuth, controllers.SendJobApplication)
		api.PUT("/jobs/:applicationid", middleware.CheckAuth, controllers.AcceptJobApplication)
	}
}
