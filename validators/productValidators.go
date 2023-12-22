package validators

import (
	"github.com/woonmapao/go-shop-api/initializer"
	"github.com/woonmapao/go-shop-api/models"
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

// Check if a product with the given name already exists in the database
func IsProductNameDuplicate(name string) bool {
	var existingProduct models.Product
	err := initializer.DB.Where("name = ?", name).First(&existingProduct).Error
	return err == nil
}
