package validators

import (
	"fmt"

	"github.com/woonmapao/go-shop-api/models"
)

// Validate order detail data
func ValidateOrderDetailData(data struct {
	OrderID   int     `json:"orderId" binding:"required"`
	ProductID int     `json:"productId" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,gte=1"`
	Subtotal  float64 `json:"subtotal"`
}) error {

	// Example: Check if the order and product exist in the database
	if !orderExists(data.OrderID) {
		return fmt.Errorf("order with ID %d does not exist", data.OrderID)
	}

	if !productExists(data.ProductID) {
		return fmt.Errorf("product with ID %d does not exist", data.ProductID)
	}

	// Check if the product has sufficient stock
	if !hasSufficientStock(data.ProductID, data.Quantity) {
		return fmt.Errorf("insufficient stock for product with ID %d", data.ProductID)
	}

	return nil
}

// Validate updated order detail data
func ValidateUpdatedOrderDetailData(data models.OrderDetail) error {
	// Add your validation logic here
	// Example: Check if the updated order detail data is valid before updating
	// Return an error if validation fails

	// Replace the example logic with your actual validation checks
	if data.Quantity <= 0 {
		return fmt.Errorf("Quantity must be greater than zero")
	}

	return nil
}
