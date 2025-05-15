package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiofisio/golang-react/controllers"
)

func AuthRoutes(router *gin.Engine) {

	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
	}

}
