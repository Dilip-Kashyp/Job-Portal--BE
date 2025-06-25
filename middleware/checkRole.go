package middleware

import (
	"net/http"

	"github.com/Dilip-Kashyp/job-portal-backend/constants"
	"github.com/gin-gonic/gin"
)

func CheckRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": constants.UNAUTHORIZED_ACCESS})
			c.Abort()
			return
		}
		c.Next()
	}
}
