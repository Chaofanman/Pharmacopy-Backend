package main

import (
	"fmt"
	"log"
	"os"

	db "github.com/chaofanman/pharmacopy-rest/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// f, err := os.Open("input_files/db3.csv")
	// utils.Check(err)
	// parser.CSVParser(f)
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	database := db.InitDB()

	fmt.Println("Database: ", database)

	router := gin.Default()
	router.Use(gin.Logger())

	fmt.Println("Port: ", port)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Version": 1.0,
		})
	})

	router.Run(":" + port)

}
