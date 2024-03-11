package controllers

import (
	"example/web-service-gin/api/data/response"
	service "example/web-service-gin/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BooksController 인터페이스 정의
type BookController interface {
	FindAll(ctx *gin.Context)
}

// BooksController 구현체
type BookControllerImpl struct {
	bookService service.BookService
}

// BooksController 생성자
func NewBooksController() BookController {
	bookService, err := service.NewBookServiceImpl()

	if err != nil {
		return nil
	}

	return &BookControllerImpl{bookService}
}

func (bc *BookControllerImpl) FindAll(ctx *gin.Context) {
	books, err := bc.bookService.GetBooks()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get books"})
		return
	}

	webResponse := response.Response{
		Code : http.StatusOK,
		Status:"success",
		Data: books,
	}

	ctx.JSON(http.StatusOK, webResponse)
}