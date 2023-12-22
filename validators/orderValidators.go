package validators

import (
	"github.com/woonmapao/go-shop-api/initializer"
	"github.com/woonmapao/go-shop-api/models"
)

// Check if an order with the given ID exists in the database
func orderExists(orderID int) bool {
	var order models.Order
	result := initializer.DB.First(&order, orderID)
	return result.RowsAffected > 0
}

// Validate order data
func ValidateOrderData(orderData models.Order) error {

	// Still out of idea how to validate

	return nil
}
