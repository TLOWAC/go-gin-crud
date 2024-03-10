package main

import (
	"example/web-service-gin/api/routes"
)

func main() {
	routes.SetupRoutes("localhost:8080")
}