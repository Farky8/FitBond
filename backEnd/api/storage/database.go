package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"github.com/joho/godotenv"
)

type EventInfo struct {
    ID	     int	`gorm:"primaryKey" json:"id"`
    Heading  string	`json:"heading"`
    About    string	`json:"about"`
    Applied  int	`json:"applied"`
    Capacity int	`json:"capacity"`
    Price    int	`json:"price"`
}

func DBSetUp() *gorm.DB {
    if err := godotenv.Load(); err != nil {
	log.Fatalf("Error loading .env file: %v", err)
    }

    user:= os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    host := os.Getenv("DB_HOST")

    if host == "" || user == "" || password == "" || dbname == "" {
        log.Fatal("missing required environment variables")
    }

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
	log.Fatalf("Failed to connect to database: %v", err)
    }

    if err := db.AutoMigrate(&EventInfo{}); err != nil {
	log.Fatalf("Failed to migrate database: %v", err)
    }

    return db;
}

