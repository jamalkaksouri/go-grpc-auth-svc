package db

import (
	"github.com/jamalkaksouri/go-grpc-auth-svc/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalln(err)
	}

	return Handler{db}
}
