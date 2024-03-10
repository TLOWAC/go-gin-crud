package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Postgresql struct {
	pg *gorm.DB
}

func NewORM() (*Postgresql, error) {

	// DB_URL 환경 변수 값 조회
	dbURL := os.Getenv("DB_URL")
	
	// DB_URL 이 없는 경우 예외처리 및 초기값 할당
	if dbURL == "" {
		dbURL = "postgresql://postgres:root12#@localhost:5432/go-local"
	}

	// gorm, postgresql 연동
	pgConn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	
	if err != nil {
		log.Fatalln(err)
	}
	// db.Migrator().AutoMigrate(&models.Book{})

	// handler DI 에서 pg orm conn 을 재사용 할 수 있도록 연결값을 담아 반환
	return &Postgresql {
		pg: pgConn,
	}, err
}