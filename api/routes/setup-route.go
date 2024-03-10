package routes

import (
	"example/web-service-gin/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(addr string) {
    router := gin.Default()
    v1 := router.Group("/")

    router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

    controllers.AddBooksRoutes(v1)
		
    router.Run()
}