package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, code int, msg string) {
	log.Printf("[Error]%s", msg)
	c.JSON(code, gin.H{"error": msg})
}
