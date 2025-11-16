package database

import (
	"log"

	"github.com/aayushchugh/ayushchugh.com-api/config/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	envConfig := env.LoadEnv()
	_, err := gorm.Open(postgres.Open(envConfig.DATABASE_URL), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	log.Println("Connected to database successfully")

}
