package database

import (
	"fmt"
	"os"

	"github.com/chaofanman/pharmacopy-rest/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //needed for postgres driver
)

var db *gorm.DB

//InitDB returns a postgres database
func InitDB() *gorm.DB {
	url := os.Getenv("DATABASE_URL") + "?sslmode=disable"
	fmt.Println("URL: ", url)
	db, err := gorm.Open("postgres", url)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.Drug{},
	)

	return db
}

//GetDB returns db
func GetDB() *gorm.DB {
	return db
}
