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
		Name:  body.Name,
		Email: body.Email,
		Age:   body.Age,
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

	// Get id off url
	id := c.Param("id")

	// Get the user

	var user models.UserDetail
	initializer.DB.Find(&user, id)

	// Respond the finds
	c.JSON(200, gin.H{
		"found user:": user,
	})

}

func GetAllUsers(c *gin.Context) {

	//Get all user
	var users []models.UserDetail
	initializer.DB.Find(&users)

	// Respond the finds
	c.JSON(200, gin.H{
		"users:": users,
	})
}

func UpdateUser(c *gin.Context) {

	// Get ID off param
	id := c.Param("id")

	// Get data of request body
	var body struct {
		Name     string
		Email    *string
		Age      uint8
		Birthday *time.Time
	}

	c.Bind(&body)

	// Get the users to update from db
	var user models.UserDetail
	initializer.DB.First(&user, id)

	// Update it
	initializer.DB.Model(&user).Updates(models.UserDetail{
		Name:  body.Name,
		Email: body.Email,
		Age:   body.Age,
	})

	// Respond the updated post
	c.JSON(200, gin.H{
		"updated user:": user,
	})

}

func DeleteUser(c *gin.Context) {

	// Get the id off url
	id := c.Param("id")

	// Delete the user
	initializer.DB.Delete(&models.UserDetail{}, id)

	// Respond
	c.JSON(200, gin.H{
		"delete:": "completed",
	})

}
