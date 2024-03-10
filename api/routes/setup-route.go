package routes

import (
	bookAPI "example/web-service-gin/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
    router := gin.Default()
    v1 := router.Group("/")

    router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

    bookAPI.AddBooksRoutes(v1)
		
    return router
}