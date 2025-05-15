package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		method := c.Request.Method
		path := c.Request.URL.Path

		c.Next()

		status := c.Writer.Status()
		duration := time.Since(start)

		fmt.Printf("Method: %s, Path: %s, Status: %d, Duration: %s\n", method, path, status, duration)
	}
}
