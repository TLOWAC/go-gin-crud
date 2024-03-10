package controllers

import (
	service "example/web-service-gin/api/services"

	"github.com/gin-gonic/gin"
)

func AddBooksRoutes(rg *gin.RouterGroup) {
	book := rg.Group("/books")
	h, _ := service.NewBookServiceImpl()
	
	book.GET("", h.GetBooks)

	// content.GET("/:id", h.GetContent)
	// content.POST("", h.AddContent)
	// content.PATCH("/:id", h.UpdateContent)
	// content.DELETE("/:id", h.DeleteContent)
}