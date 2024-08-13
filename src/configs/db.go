package configs

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbOnce sync.Once

func InitDB() {
	dbOnce.Do(func() {
		url := os.Getenv("URL")
		var err error
		DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
	})
}
