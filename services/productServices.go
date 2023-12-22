package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/woonmapao/go-shop-api/initializer"
	"github.com/woonmapao/go-shop-api/models"
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

	// Extract filter parameters
	priceMin := c.Query("priceMin")
	priceMax := c.Query("priceMax")
	category := c.Query("category")

	// Start building the database query
	query := initializer.DB.Where("name LIKE ?", "%"+searchCriteria+"%")

	// Add filters to the query if provided
	if priceMin != "" {
		query = query.Where("price >= ?", priceMin)
	}

	if priceMax != "" {
		query = query.Where("price <= ?", priceMax)
	}

	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Execute the database query
	var searchResults []models.Product
	err := query.Find(&searchResults).Error
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
