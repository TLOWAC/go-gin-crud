package handler

import (
	"example/web-service-gin/api/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BookService 에 정의할 함수 인터페이스 정의
type BookService interface {
	GetBooks(c *gin.Context)
	// GetBook(c *gin.Context)
	// AddBook(c *gin.Context)
	// UpdateBook(c *gin.Context)
	// DeleteBook(c *gin.Context)
}

type BookServiceImpl struct {
	bookRepository repositories.BookRepository
}

func NewBookServiceImpl() (BookService, error) {

	db, err := repositories.NewBookRepositoryImpl()
	if err != nil {
		return nil, err
	}

	return &BookServiceImpl{bookRepository: db}, nil
}


func (h *BookServiceImpl) GetBooks(c *gin.Context) {
	contents, err := h.bookRepository.GetAllContents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contents)
}