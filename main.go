package main

import (
	"github.com/Dilip-Kashyp/job-portal-backend/config"
	"github.com/Dilip-Kashyp/job-portal-backend/models"
	"github.com/Dilip-Kashyp/job-portal-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Job{},
		&models.Recruiter{},
	)

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
