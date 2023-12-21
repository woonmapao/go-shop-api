package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/models"
)

func AddUser(c *gin.Context) {
	// Handle the creation of a new user

	// Get data from the request body
	var body struct {
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create user in the database
	user := models.User{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
	}

	if err := initializer.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Return status
	c.JSON(http.StatusOK, gin.H{
		"createdUser": user,
	})

}

func GetUserByID(c *gin.Context) {
	// Retrieve a specific user based on their ID

	// Get ID from URL param
	id := c.Param("id")

	// Get the user from the database
	var user models.User
	err := initializer.DB.First(&user, id).Error
	if err != nil {
		// Handle user not found or other errors
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	// Respond with the found user
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func GetAllUsers(c *gin.Context) {
	// Fetch a list of all users from the database

	// Get all users from the database
	var users []models.User
	err := initializer.DB.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
		})
		return
	}

	// Respond with the found users
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})

}

func UpdateUser(c *gin.Context) {
	// Handle the update of an existing user

	// Get ID from URL param
	id := c.Param("id")

	// Get data from request body
	var body struct {
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check if the user with the given ID exists
	var user models.User
	err = initializer.DB.First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	// Update user
	initializer.DB.Model(&user).Updates(models.User{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
	})

	// Respond with the updated user
	c.JSON(http.StatusOK, gin.H{
		"updatedUser": user,
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
