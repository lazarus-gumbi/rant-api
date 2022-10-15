package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lazarus-gumbi/sample_api/initializers"
	"github.com/lazarus-gumbi/sample_api/models"
)

func RantCreate(c *gin.Context) {

	var body struct {
		Body    string
		Title   string
		Company string
	}

	c.Bind(&body)

	rant := models.Post{Title: body.Title, Company: body.Company, Body: body.Body}
	result := initializers.DB.Create(&rant)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"rant": rant,
	})
}

func RantsIndex(c *gin.Context) {
	var rants []models.Post
	initializers.DB.Find(&rants)

	c.JSON(200, gin.H{
		"rants": rants,
	})
}

func RantsShow(c *gin.Context) {
	company := c.Param("company")

	var rants []models.Post
	initializers.DB.Where("Company LIKE ?", company).Find(&rants)

	c.JSON(200, gin.H{
		"rant": rants,
	})
}

func RantDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
