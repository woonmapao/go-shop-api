package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/models"
	"github.com/woonmapao/user-management/validators"
)

func GetAllOrders(c *gin.Context) {
	// Fetch a list of all orders from the database

	// Retrieve orders from the database
	var orders []models.Order
	err := initializer.DB.Find(&orders).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch orders",
		})
		return
	}

	// Return a JSON response with the list of orders
	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func GetOrderByID(c *gin.Context) {
	// Extract order ID from the request parameters
	orderID := c.Param("id")

	// Query the database for the order with the specified ID
	var order models.Order
	err := initializer.DB.First(&order, orderID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Order not found",
		})
		return
	}

	// Return a JSON response with the order details
	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

func CreateOrder(c *gin.Context) {
	// Extract order data from the request body
	var orderData models.Order
	err := c.ShouldBindJSON(&orderData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validate the input data
	err = validators.ValidateOrderData(orderData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create a new order in the database
	err = initializer.DB.Create(&orderData).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create order",
		})
		return
	}

	// Return a JSON response with the newly created order
	c.JSON(http.StatusCreated, gin.H{
		"order": orderData,
	})
}

// UpdateOrder handles the update of an existing order
func UpdateOrder(c *gin.Context) {
	// Extract order ID from the request parameters
	orderID := c.Param("id")

	// Extract updated order data from the request body
	var updatedOrderData models.Order
	err := c.ShouldBindJSON(&updatedOrderData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validate the input data
	err = validators.ValidateOrderData(updatedOrderData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Get the existing order from the database
	var existingOrder models.Order
	err = initializer.DB.First(&existingOrder, orderID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Order not found",
		})
		return
	}

	// Update the order in the database
	initializer.DB.Model(&existingOrder).Updates(updatedOrderData)

	// Return a JSON response with the updated order
	c.JSON(http.StatusOK, gin.H{
		"updated_order": existingOrder,
	})
}

func DeleteOrder(c *gin.Context) {
	// Delete an order based on its ID

	// Extract order ID from the request parameters
	// Delete the order from the database
	// Return a JSON response indicating success or failure
}

func GetUserOrders(c *gin.Context) {
	// Fetch all orders associated with a specific user

	// Extract user ID from the request parameters
	// Query the database for orders associated with the user
	// Return a JSON response with the user's orders
}

func GetOrderDetails(c *gin.Context) {
	// Fetch all details (products) associated with a specific order

	// Extract order ID from the request parameters
	// Query the database for details (products) associated with the order
	// Return a JSON response with the order details
}
