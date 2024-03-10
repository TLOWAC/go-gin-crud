package db

import (
	"example/web-service-gin/api/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Postgresql struct {
	*gorm.DB
}

func NewPgManager() (*Postgresql, error) {

	// DB_URL 환경 변수 값 조회
	dbURL := os.Getenv("DB_URL")
	
	// DB_URL 이 없는 경우 예외처리 및 초기값 할당
	if dbURL == "" {
		dbURL = "postgres://root:root123@localhost:5432/go-local"
	}

	// gorm, postgresql 연동
	pgConn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	
	if err != nil {
		log.Fatalln(err)
	}
	pgConn.Migrator().AutoMigrate(&models.Book{})

	// handler DI 에서 pg orm conn 을 재사용 할 수 있도록 연결값을 담아 반환
	return &Postgresql {
		pgConn,
	}, err
}
