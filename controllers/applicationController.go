package controllers

import (
	"net/http"

	"github.com/Dilip-Kashyp/job-portal-backend/config"
	"github.com/Dilip-Kashyp/job-portal-backend/models"
	"github.com/gin-gonic/gin"
)

func ApplyJob(c *gin.Context) {
	var data struct {
		JobID uint `json:"job_id"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	application := models.Application{
		UserID: userID.(uint),
		JobID:  data.JobID,
		Status: "Applied",
	}
	config.DB.Create(&application)
	c.JSON(http.StatusCreated, gin.H{"message": "Job applied"})
}

func GetAppliedJob(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var apps []models.Application
	config.DB.Preload("Job").Where("user_id = ?", userID).Find(&apps)
	c.JSON(http.StatusOK, apps)
}

func UpdateJobStatus(c *gin.Context) {
	var input struct {
		JobID  uint   `json:"job_id"`
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Model(&models.Job{}).Where("id = ?", input.JobID).Update("status", input.Status)
	c.JSON(http.StatusOK, gin.H{"message": "Job status updated"})
}

func GetApplicants(c *gin.Context) {
	var apps []models.Application
	config.DB.Preload("User").Preload("Job").Find(&apps)
	c.JSON(http.StatusOK, apps)
}
