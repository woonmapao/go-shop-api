package services

import (
	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/models"
)

// Update stock in the database
func updateStock(productID, quantity int) error {
	var product models.Product
	result := initializer.DB.First(&product, productID)
	if result.Error != nil {
		return result.Error
	}

	newStock := product.StockQuantity - quantity
	initializer.DB.Model(&product).Update("StockQuantity", newStock)

	return nil
}
