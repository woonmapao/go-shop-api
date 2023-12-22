package validators

import (
	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/models"
)

// Check if a product with the given ID exists in the database
func productExists(productID int) bool {
	var product models.Product
	result := initializer.DB.First(&product, productID)
	return result.RowsAffected > 0
}

// Check if a product has sufficient stock for the given quantity
func hasSufficientStock(productID, quantity int) bool {
	var product models.Product
	result := initializer.DB.First(&product, productID)
	return result.RowsAffected > 0 && product.StockQuantity >= quantity
}
