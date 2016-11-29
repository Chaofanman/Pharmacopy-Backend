package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chaofanman/pharmacopy-rest/controllers"
	db "github.com/chaofanman/pharmacopy-rest/database"
	"github.com/chaofanman/pharmacopy-rest/parser"
	"github.com/chaofanman/pharmacopy-rest/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	f, err := os.Open("input_files/db3.csv")
	utils.Check(err)
	parser.CSVParser(f)
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	database := db.Init()

	fmt.Println("Database: ", database)

	router := gin.Default()
	router.Use(gin.Logger())

	fmt.Println("Port: ", port)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Version": 1.0,
		})
	})

	drugs := new(controllers.DrugController)

	v1 := router.Group("api/v1")
	{
		v1.GET("/drugs", drugs.GetDrugs)
	}

	router.Run(":" + port)
}
