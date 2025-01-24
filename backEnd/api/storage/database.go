package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type EventInfo struct {
    ID	     string	`gorm:"primaryKey" json:"id,omitempty"`
    Heading  string	`json:"heading"`
    About     string	`json:"about"`
    Applied  int	`json:"applied,omitempty"`
    Capacity int	`json:"capacity"`
    Price    int	`json:"price"`
}

func DBSetUp() *gorm.DB {
    url := os.Getenv("DB_URL")
    if url == "" {
	log.Fatalf("DB_DSN environment variable is not set")
    }

    db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
    if err != nil {
	log.Fatalf("Failed to connect to database: %v", err)
    }

    if err := db.AutoMigrate(&EventInfo{}); err != nil {
	log.Fatalf("Failed to migrate database: %v", err)
    }

    return db;
}

