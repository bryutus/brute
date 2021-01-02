package infrastructure

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	db  *gorm.DB
	err error
)

func Init() *gorm.DB {
	config := getConfig()

	db, err = gorm.Open("postgres", config)
	if err != nil {
		log.Fatalf("Error db init: %v", err)
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}

func getConfig() string {
	err = godotenv.Load()
	if err != nil {
		log.Printf("Alert loading .env: %v", err)
	}

	return os.Getenv("DB_CONNECTION")
}
