package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/models"
)

func AddUser(c *gin.Context) {

	// Get data off req body

	var body struct {
		Name     string
		Email    *string
		Age      uint8
		Birthday *time.Time
	}

	c.Bind(&body)

	// Add user to db

	user := models.UserDetail{
		Name:     body.Name,
		Email:    body.Email,
		Age:      body.Age,
		Birthday: body.Birthday,
	}

	result := initializer.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return status

	c.JSON(200, gin.H{
		"created :": user,
	})

}

func GetUserByID(c *gin.Context) {

}

func GetAllUsers(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
