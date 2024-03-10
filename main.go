package main

import (
	"example/web-service-gin/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
    v1 := router.Group("/")

    router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		routes.AddBooksRoutes(v1)
		
    router.Run("localhost:8080")
}