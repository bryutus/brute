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

func Init(filenames ...string) *gorm.DB {
	config := getConfig(filenames)

	db, err = gorm.Open("postgres", config)
	if err != nil {
		log.Fatalf("Error db init: %v", err)
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}

func getConfig(filenames []string) string {
	err = godotenv.Load(env(filenames))
	if err != nil {
		log.Printf("Alert loading .env: %v", err)
	}

	return os.Getenv("DB_CONNECTION")
}

func env(files []string) string {
	env := ".env"
	for _, f := range files {
		env = f
	}
	return env
}
