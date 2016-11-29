package controllers

import (
	"fmt"

	"github.com/chaofanman/pharmacopy-rest/database"
	"github.com/chaofanman/pharmacopy-rest/models"
	"github.com/gin-gonic/gin"
)

//DrugController struct
type DrugController struct{}

//GetDrugs queries all drugs
func (ctrl DrugController) GetDrugs(c *gin.Context) {
	fmt.Println("AJSDFJASDLFJAL;SKJF;LDSJFLS")
	var drugs []models.Drug

	db := database.Init()

	c.JSON(200, "Here")
	err := db.Find(&drugs).Error

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"Error": "Error finding drugs",
		})
	}

	c.JSON(200, drugs)
}
