package controllers

import (
	"net/http"

	"github.com/Dilip-Kashyp/job-portal-backend/config"
	"github.com/Dilip-Kashyp/job-portal-backend/constants"
	"github.com/Dilip-Kashyp/job-portal-backend/models"
	"github.com/Dilip-Kashyp/job-portal-backend/payload"
	"github.com/Dilip-Kashyp/job-portal-backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user payload.RegisterPayload
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.INVALID_ERROR_MESSAGE})
		return
	}
	hashedPassword := utils.HashPassword(user.Password)
	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
		Role:     user.Role,
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.DATABSE_INVALID_SERVER_MESSAGE})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": constants.DATABASE_SUCCESS_MESSAGE})
}

func LoginUser(c *gin.Context) {
	var input payload.LoginPayload
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.INVALID_ERROR_MESSAGE})
		return
	}
	var user models.User
	config.DB.Where("email = ?", input.Email).First(&user)
	if user.ID == 0 || !utils.CheckPassword(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.INVALID_EMAIL_PASSWORD_MESSAGE})
		return
	}
	token, _ := utils.GenerateJWT(user.ID, user.Role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetUsers(c *gin.Context) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.USER_NOT_FOUND})
		return
	}

	userID, ok := userIDValue.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.INVALID_SERVER_MESSAGE})
		return
	}

	var user models.User
	if err := config.DB.First(&user, uint(userID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": constants.USER_NOT_FOUND})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateProfile(c *gin.Context) {
	userID := c.GetInt("userID")
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.INVALID_ERROR_MESSAGE})
		return
	}
	config.DB.Model(&models.User{}).Where("id = ?", userID).Updates(input)
	c.JSON(http.StatusOK, gin.H{"message": constants.PROFILE_UPDATED_MESSAGE})
}
