package infrastructure

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() *gorm.DB {
	config := "host=db user=admin password=admin dbname=app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err = gorm.Open("postgres", config)
	if err != nil {
		fmt.Println("db init error: ", err)
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}
