package main

import (
	"log"
	"os"

	"github.com/chaofanman/pharmacopy-rest/controllers"
	"github.com/chaofanman/pharmacopy-rest/parser"
	"github.com/gin-gonic/gin"
)

func main() {
	parser.ParseFile("input_files/db1.csv")
	parser.ParseFile("input_files/db2.csv")
	parser.ParseFile("input_files/db3.csv")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Version": 1.0,
		})
	})

	drugs := new(controllers.DrugController)

	v1 := router.Group("api/v1")
	{
		v1.POST("/drugs", drugs.GetDrugs)
		v1.POST("/drugs/:name", drugs.GetDrug)
	}

	router.Run(":" + port)
}
