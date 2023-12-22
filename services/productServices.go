package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/models"
)

// Update stock in the database
func UpdateStock(productID, quantity int) error {
	var product models.Product
	result := initializer.DB.First(&product, productID)
	if result.Error != nil {
		return result.Error
	}

	newStock := product.StockQuantity - quantity
	initializer.DB.Model(&product).Update("StockQuantity", newStock)

	return nil
}

func SearchProducts(c *gin.Context) {
	// Extract search criteria from the request parameters or query string
	searchCriteria := c.Query("searchCriteria")

	// Query the database for products that match the criteria
	var searchResults []models.Product
	err := initializer.DB.Where("name LIKE ?", "%"+searchCriteria+"%").Find(&searchResults).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch search results",
		})
		return
	}

	// Return a JSON response with the search results
	c.JSON(http.StatusOK, gin.H{
		"searchResults": searchResults,
	})
}
