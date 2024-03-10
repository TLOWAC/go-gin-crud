package main

import (
	"github.com/gin-gonic/gin"

	"example/web-service-gin/api/routes"
)

func main() {
	
	router := gin.Default()

	routes.SetupRoutes()
	
	router.Run("localhost:8080")
}