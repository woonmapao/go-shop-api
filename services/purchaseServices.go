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
	// Extract user and product information
	var purchaseData struct {
		UserID   int                      `json:"userId"`
		Products []models.ProductPurchase `json:"products"`
	}

	err := c.ShouldBindJSON(&purchaseData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validate input data
	err = validators.ValidatePurchaseData(purchaseData.UserID, purchaseData.Products)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Update stock and create order details
	var orderDetails []models.OrderDetail
	var totalAmount float64

	for _, product := range purchaseData.Products {
		// Update stock
		err := UpdateStock(product.ProductID, product.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update stock",
			})
			return
		}

		// Create order detail
		orderDetail := models.OrderDetail{
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
			Subtotal:  calculateSubtotal(product.ProductID, product.Quantity),
		}

		orderDetails = append(orderDetails, orderDetail)
		totalAmount += orderDetail.Subtotal
	}

	// Create order
	order := models.Order{
		UserID:      purchaseData.UserID,
		OrderDate:   time.Now(),
		TotalAmount: totalAmount,
		Status:      "Pending", // Set initial status as "Pending"
	}

	err = initializer.DB.Create(&order).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create order",
		})
		return
	}

	// Link order details to the created order
	for _, orderDetail := range orderDetails {
		orderDetail.OrderID = int(order.ID)
		err = initializer.DB.Create(&orderDetail).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create order detail",
			})
			return
		}
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"message":      "Purchase successful",
		"orderID":      order.ID,
		"totalAmount":  totalAmount,
		"orderDetails": orderDetails,
	})
}

// Function to calculate subtotal
func calculateSubtotal(productID, quantity int) float64 {
	// Fetch product price from the database
	var product models.Product
	err := initializer.DB.First(&product, productID).Error
	if err != nil {
		return 0.0
	}

	// Calculate subtotal based on product price and quantity
	subtotal := float64(quantity) * product.Price
	return subtotal
}
