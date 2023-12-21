package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/models"
)

func AddUser(c *gin.Context) {
	// Handle the creation of a new user

	// Get data off req body

	var body struct {
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
	}

	c.Bind(&body)

	// Add user to db

	user := models.User{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
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
	// Retrieve a specific user based on their ID

	// Get id off url
	id := c.Param("id")

	// Get the user

	var user models.User
	initializer.DB.Find(&user, id)

	// Respond the finds
	c.JSON(200, gin.H{
		"found user:": user,
	})

}

func GetAllUsers(c *gin.Context) {
	// Fetch a list of all users from the database

	//Get all user
	var users []models.User
	initializer.DB.Find(&users)

	// Respond the finds
	c.JSON(200, gin.H{
		"users:": users,
	})
}

func UpdateUser(c *gin.Context) {
	// Handle the update of an existing user

	// Get ID off param
	id := c.Param("id")

	// Get data of request body
	var body struct {
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
	}

	c.Bind(&body)

	// Get the users to update from db
	var user models.User
	initializer.DB.First(&user, id)

	// Update it
	initializer.DB.Model(&user).Updates(models.User{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
	})

	// Respond the updated post
	c.JSON(200, gin.H{
		"updated user:": user,
	})

}

// func DeleteUser(c *gin.Context) {
// 	// Delete a user based on their ID

// 	// Get the id off url
// 	id := c.Param("id")

// 	// Delete the user
// 	initializer.DB.Delete(&models.User{}, id)

// 	// Respond
// 	c.JSON(200, gin.H{
// 		"delete:": "completed",
// 	})

// }
