package infrastructure

import (
	"fmt"
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
		log.Fatalf("Error db init: ", err)
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}

func getConfig() string {
	err = godotenv.Load()
	if err != nil {
		log.Printf("Alert loading .env:", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_DATABASE")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tokyo", host, port, user, pass, name)
}
