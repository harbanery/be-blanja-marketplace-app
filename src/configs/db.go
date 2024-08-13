package configs

import (
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbOnce sync.Once

func InitDB() {
	dbOnce.Do(func() {
		url := os.Getenv("URL")

		var err error
		DB, err = gorm.Open(postgres.Open(url), &gorm.Config{
			PrepareStmt: true,
		})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	})
}
