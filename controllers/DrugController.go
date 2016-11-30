package controllers

import (
	"fmt"
	"sort"

	"github.com/chaofanman/pharmacopy-rest/database"
	"github.com/chaofanman/pharmacopy-rest/models"
	"github.com/gin-gonic/gin"
)

//DrugController struct
type DrugController struct{}

//GetDrugs queries all drugs
func (ctrl DrugController) GetDrugs(c *gin.Context) {
	var drugs models.Drugs

	db := database.Init()

	letter := c.Query("letter")
	fmt.Println("Letter: ", letter)

	if letter == "" {
		err := db.Table("drugs").Select("id, name").Scan(&drugs).Error
		if err != nil {
			c.JSON(400, gin.H{
				"Error": "Error finding all drugs",
			})
		}
	} else {
		err := db.Where("name LIKE ?", letter+"%").Select("id, name").Find(&drugs).Error
		if err != nil {
			c.JSON(400, gin.H{
				"Error": "Error finding drugs with letter " + letter,
			})
		}
	}

	sort.Sort(drugs)

	c.JSON(200, drugs)
}
