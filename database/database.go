package database

import (
	"os"

	"github.com/chaofanman/pharmacopy-rest/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //needed for postgres driver
)

var db *gorm.DB

//Init returns a postgres database
func Init() *gorm.DB {
	url := os.Getenv("DATABASE_URL") + "?sslmode=disable"
	db, err := gorm.Open("postgres", url)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.Drug{},
	)

	return db
}
