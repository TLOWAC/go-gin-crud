package main

import (
	"example/web-service-gin/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 기본 router 셋팅
  router := gin.Default()

	// /api/v1 url prefix 추가
	v1Router := router.Group("/api/v1")

	// 테스트용 ping 엔드포인트
	v1Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// ctrl 붙임
	routes.AddBooksRoutes(v1Router)
	
	// 서버 동작
	router.Run("localhost:8080")
}