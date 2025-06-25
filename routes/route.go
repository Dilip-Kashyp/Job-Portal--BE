package routes

import (
	"github.com/Dilip-Kashyp/job-portal-backend/controllers"
	"github.com/Dilip-Kashyp/job-portal-backend/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Public Routes
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "API is running!"})
	})
	api.POST("/user/register", controllers.CreateUser)
	api.POST("/user/login", controllers.LoginUser)
	api.GET("/job/get-all-jobs", controllers.GetJobs)

	// Protected Routes - Auth Required
	auth := api.Group("/")
	auth.Use(middleware.IsAuthenticated())
	{
		auth.GET("/user/get-current-user", controllers.GetUsers)
		auth.POST("/user/update-profile", controllers.UpdateProfile)
		auth.GET("/job/get-job:id", controllers.GetJobsByID)
		auth.GET("/user/get-profile:id", controllers.GetJobsByID)

		// Role based auth
		auth.POST("/job/apply-job", middleware.CheckRole("student"), controllers.ApplyJob)
		// auth.POST("/job-suggestion", middleware.CheckRole("student"), controllers.JobSuggestion)
		auth.POST("/job/get-applied-job", middleware.CheckRole("student"), controllers.GetAppliedJob)

		auth.PATCH("/job/update-job", middleware.CheckRole("recruiter"), controllers.UpdateJobStatus)
		auth.POST("/job/create-job", middleware.CheckRole("recruiter"), controllers.CreateJob)
		auth.GET("/job/get-applied-job", middleware.CheckRole("recruiter"), controllers.GetApplicants)
		auth.GET("/get-student-suggestion", middleware.CheckRole("recruiter"), controllers.GetApplicants)
	}
}
