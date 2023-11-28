package database

import (
	"finalproject/models"
	"fmt"

	"os"

	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
	)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to db", err)
	}
	db.Debug().AutoMigrate(models.User{}, models.Comment{}, models.UserPhoto{}, models.UserSocialMedia{})

}
func GetDB() *gorm.DB {
	return db
}
