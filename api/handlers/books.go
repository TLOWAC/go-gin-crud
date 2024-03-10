package handler

import (
	"example/web-service-gin/api/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	bookRepository repositories.BookRepository
}


type BookHandlerInterface interface {
	GetBooks(c *gin.Context)
	// GetBook(c *gin.Context)
	// AddBook(c *gin.Context)
	// UpdateBook(c *gin.Context)
	// DeleteBook(c *gin.Context)
}

func (h *Handler) GetBooks(c *gin.Context) {
	contents, err := h.bookRepository.GetAllContents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contents)
}


// BookHandlerInterface 생성자
func NewHandler() (BookHandlerInterface, error) {

	db, err := repositories.NewBookORM()
	if err != nil {
		return nil, err
	}

	return &Handler{
		bookRepository: db,
	}, nil
}