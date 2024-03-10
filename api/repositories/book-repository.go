package repositories

import (
	"example/web-service-gin/api/models"
	"example/web-service-gin/internal/db"
	"log"
)

// BookRepo 관련 함수 인터페이스 정의
type BookRepository interface {
	GetAllContents() ([]models.Book, error)
	// GetAllContent() ([]models.Book, error)
}

// BookRepo 인터페이스 함수 구현
func (pg *BookORM) GetAllContents() ([]models.Book, error) {	
	var books []models.Book
	err := pg.Find(&books).Error
	return books, err
}

// BookORM 함수 작성시 DI 형태로 주입할 DB Conn 정의
type BookORM struct {
	*db.Postgresql
}

// DB 연결
func NewBookORM() (*BookORM, error) {
	pg, err := db.NewPgManager()

	if err != nil {
		log.Fatalln(err)
	}

	return &BookORM{
		pg,
	}, err
}

