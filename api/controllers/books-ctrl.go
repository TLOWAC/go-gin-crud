package controllers

import (
	handler "example/web-service-gin/api/handlers"

	"github.com/gin-gonic/gin"
)

func AddBooksRoutes(rg *gin.RouterGroup) {
	book := rg.Group("/books")
	h, _ := handler.NewHandler()
	
	book.GET("", h.GetBooks)

	// content.GET("/:id", h.GetContent)
	// content.POST("", h.AddContent)
	// content.PATCH("/:id", h.UpdateContent)
	// content.DELETE("/:id", h.DeleteContent)
}