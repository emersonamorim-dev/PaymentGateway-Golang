package utils

import (
	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
	c.Abort()
}

func RespondWithSuccess(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"message": message})
}
