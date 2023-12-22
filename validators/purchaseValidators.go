package validators

import "fmt"

// Validate purchase data
func ValidatePurchaseData(data struct{ UserID, ProductID, Quantity int }) error {

	// Check if the user and product exist in the database
	if !userExists(data.UserID) {
		return fmt.Errorf("user with ID %d does not exist", data.UserID)
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
