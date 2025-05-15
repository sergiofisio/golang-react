package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SendErrorResponse(c *gin.Context, statusCode int, message string, function string) {
	log.Printf("[%s] ❌ Erro: %s", function, message)
	c.JSON(statusCode, gin.H{"error": message})
}
