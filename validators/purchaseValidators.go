package validators

import (
	"errors"

	"github.com/woonmapao/user-management/models"
)

// ValidatePurchaseData validates the provided purchase data
func ValidatePurchaseData(userId int, products []models.ProductPurchase) error {
	// Check if UserID is valid
	if userId <= 0 {
		return errors.New("UserID must be greater than zero")
	}

	// Check if there are products in the purchase
	if len(products) == 0 {
		return errors.New("At least one product must be included in the purchase")
	}

	// Validate each product in the purchase
	for _, product := range products {
		// Check if ProductID is valid
		if product.ProductID <= 0 {
			return errors.New("ProductID must be greater than zero")
		}

		// Check if Quantity is valid
		if product.Quantity <= 0 {
			return errors.New("Quantity must be greater than zero")
		}
	}

	return nil
}
