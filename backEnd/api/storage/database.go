package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
	"log"
	"os"
)

type TrainInfo struct {
    ID	     string	`gorm:"primaryKey" json:"id,omitempty"`
    Heading  string	`json:"heading"`
    Body     string	`json:"body"`
    Applied  int	`json:"applied"`
    Capacity int	`json:"capacity"`
    Price    int	`json:"price"`
}

type Trainings struct {
    DB *gorm.DB
}

func DBSetUp() *Trainings {
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
	log.Fatalf("DB_DSN environment variable is not set")
    }

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    //Logger: logger.Default.LogMode(logger.Info),
})
    if err != nil {
	log.Fatalf("Failed to connect to database: %v", err)
    }

    if err := db.AutoMigrate(&TrainInfo{}); err != nil {
	log.Fatalf("Failed to migrate database: %v", err)
    }

    return &Trainings{
	DB: db,
    }
}

