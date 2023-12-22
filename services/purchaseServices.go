package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/models"
	"github.com/woonmapao/user-management/validators"
)

func PurchaseProduct(c *gin.Context) {
	// Extract user and product information from the request
	var purchaseData struct {
		UserID    int `json:"userId" binding:"required"`
		ProductID int `json:"productId" binding:"required"`
		Quantity  int `json:"quantity" binding:"required,gte=1"`
	}

	err := c.ShouldBindJSON(&purchaseData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	purchase := struct {
		UserID    int
		ProductID int
		Quantity  int
	}{
		UserID:    purchaseData.UserID,
		ProductID: purchaseData.ProductID,
		Quantity:  purchaseData.Quantity,
	}

	// Validate input data
	err = validators.ValidatePurchaseData(purchase)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Handle the purchase (update stock, create order)
	err = updateStockAndCreateOrder(purchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process the purchase",
		})
		return
	}

	// Return a JSON response indicating the success of the purchase
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

// Update stock and create an order in the database
func updateStockAndCreateOrder(data struct{ UserID, ProductID, Quantity int }) error {

	// Update stock
	err := updateStock(data.ProductID, data.Quantity)
	if err != nil {
		return err
	}

	// Create order detail and order records
	err = createOrderAndDetail(data)
	if err != nil {
		return err
	}

	return nil
}

// Create order and order detail records in the database
func createOrderAndDetail(data struct{ UserID, ProductID, Quantity int }) error {

	subtotal := calculateSubtotal(data.ProductID, data.Quantity)

	// Create order
	order := models.Order{
		UserID:      data.UserID,
		OrderDate:   time.Now(),
		TotalAmount: subtotal,
		Status:      "Pending", // or any initial status
	}
	initializer.DB.Create(&order)

	// Create order detail
	orderDetail := models.OrderDetail{
		OrderID:   int(order.ID),
		ProductID: data.ProductID,
		Quantity:  data.Quantity,
		Subtotal:  subtotal,
	}
	initializer.DB.Create(&orderDetail)

	return nil
}

func calculateSubtotal(productID, quantity int) float64 {
	var product models.Product
	result := initializer.DB.First(&product, productID)
	if result.Error != nil {
		return 0
	}

	return float64(quantity) * product.Price
}
