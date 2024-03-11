package repositories

import (
	"example/web-service-gin/api/models"
	"example/web-service-gin/internal/db"
)

// BookRepo 관련 함수 인터페이스 정의
type BookRepository interface {
	GetAllContents() ([]models.Book, error) 
	// GetAllContent() ([]models.Book, error)
}

// BookRepositoryImpl 함수 작성시 DI 형태로 주입할 DB Conn 정의
type BookRepositoryImpl struct {
	*db.Postgresql
}

// DB 연결
func NewBookRepositoryImpl() (*BookRepositoryImpl, error) {
	pg, err := db.NewPgManager()

	if err != nil {
		return nil, err
	}

	return &BookRepositoryImpl{pg}, err
}


// BookRepo 인터페이스 함수 구현
func (pg *BookRepositoryImpl) GetAllContents() ([]models.Book, error) {	
	var books []models.Book
	err := pg.Find(&books).Error

	if err != nil {
		return nil, err
	}

	return books, nil
}