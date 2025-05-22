package database

import (
	"fmt"
	"go-test/models"
	"log"
    "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    var err error
    host := os.Getenv("POSTGRES_HOST")
    username := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    databaseName := os.Getenv("POSTGRES_DATABASE")
    port := os.Getenv("POSTGRES_PORT")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Douala", host, username, password, databaseName, port)

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    fmt.Println("Database connection established")

    DB.AutoMigrate(&models.Book{}, &models.Author{})   
}
