package controllers

import (
	"net/http"

	"github.com/Dilip-Kashyp/job-portal-backend/config"
	"github.com/Dilip-Kashyp/job-portal-backend/constants"
	"github.com/Dilip-Kashyp/job-portal-backend/models"
	"github.com/gin-gonic/gin"
)

// CreateJob handles the creation of a new job posting
func CreateJob(c *gin.Context) {
	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.INVALID_ERROR_MESSAGE})
		return
	}

	if err := config.DB.Create(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.DATABSE_INVALID_SERVER_MESSAGE})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": constants.JOB_CREATED_MESSAGE})
}

// GetJobs retrieves all job postings
func GetJobs(c *gin.Context) {
	var jobs []models.Job
	if err := config.DB.Find(&jobs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch jobs"})
		return
	}
	c.JSON(http.StatusOK, jobs)
}

// GetJobsByID fetches a single job posting by ID
func GetJobsByID(c *gin.Context) {
	id := c.Param("id")

	var job models.Job
	if err := config.DB.First(&job, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	c.JSON(http.StatusOK, job)
}
