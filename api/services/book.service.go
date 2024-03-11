package services

import (
	"example/web-service-gin/api/models"
	"example/web-service-gin/api/repository"
)

// BookService 에 정의할 함수 인터페이스 정의
type BookService interface {
	GetBooks() ([]models.Book, error) 
	// GetBook(c *gin.Context)
	// AddBook(c *gin.Context)
	// UpdateBook(c *gin.Context)
	// DeleteBook(c *gin.Context)
}

// BookService 구현체
type BookServiceImpl struct {
	bookRepository repository.BookRepository
}

// BookService 생성자
func NewBookServiceImpl() (BookService, error) {
	bookRepository, err := repository.NewBookRepositoryImpl()

	if err != nil {
		return nil, err
	}

	return &BookServiceImpl{bookRepository}, nil
}


func (bs *BookServiceImpl) GetBooks() ([]models.Book, error) {
	contents, err := bs.bookRepository.GetAllContents()

	if err != nil {
		return nil, err
	}

	return contents, nil
}