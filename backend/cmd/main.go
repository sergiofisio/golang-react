package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/sergiofisio/golang-react/config"
	"github.com/sergiofisio/golang-react/middlewares"
	"github.com/sergiofisio/golang-react/routes"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	router := gin.Default()

	router.Use(middlewares.RequestLogger())

	routes.AuthRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"init":    true,
			"status":  "ok",
			"message": "Hello World",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		start := time.Now()

		duration := time.Since(start)

		c.JSON(200, gin.H{
			"message":      "pong2",
			"duration":     duration,
			"milliseconds": duration.Milliseconds(),
		})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	host := os.Getenv("HOST")

	if host == "" {
		host = "localhost"
	}

	fmt.Printf("🚀 Servidor rodando em: http://%s:%s\n", host, port)

	err = router.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
