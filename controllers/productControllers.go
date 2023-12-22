package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/models"
	"github.com/woonmapao/user-management/validators"
)

func GetAllProducts(c *gin.Context) {
	// Retrieve products from the database
	var products []models.Product
	err := initializer.DB.Find(&products).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch products",
		})
		return
	}

	// Return a JSON response with the list of products
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func GetProductByID(c *gin.Context) {
	// Extract product ID from the request parameters
	id := c.Param("id")

	// Query the database for the product with the specified ID
	var product models.Product
	err := initializer.DB.First(&product, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	// Return a JSON response with the product details
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func CreateProduct(c *gin.Context) {
	// Extract product data from the request body
	var body struct {
		Name          string  `json:"name" binding:"required"`
		Category      string  `json:"category" binding:"required"`
		Price         float64 `json:"price" binding:"required,gt=0"`
		Description   string  `json:"description"`
		StockQuantity int     `json:"stockQuantity" binding:"required,gte=0"`
		ReorderLevel  int     `json:"reorderLevel" binding:"required,gte=0,ltfield=StockQuantity"`
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validate the input data

	// Check if the product name is unique
	isDuplicate := validators.IsProductNameDuplicate(body.Name)
	if isDuplicate {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product with this name already exists",
		})
		return
	}

	// Check if stock quantity is greater than or equal to reorder level
	if body.StockQuantity < body.ReorderLevel {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Stock quantity must be greater than or equal to reorder level",
		})
		return
	}

	// Create a new product in the database
	newProduct := models.Product{
		Name:          body.Name,
		Category:      body.Category,
		Price:         body.Price,
		Description:   body.Description,
		StockQuantity: body.StockQuantity,
		ReorderLevel:  body.ReorderLevel,
	}

	err = initializer.DB.Create(&newProduct).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create product",
		})
		return
	}

	// Return a JSON response with the newly created product
	c.JSON(http.StatusCreated, gin.H{
		"createdProduct": newProduct,
	})
}

func UpdateProduct(c *gin.Context) {
	// Extract product ID from the request parameters
	id := c.Param("id")

	// Extract updated product data from the request body
	var updatedProductData struct {
		Name          string  `json:"name" binding:"required"`
		Category      string  `json:"category" binding:"required"`
		Price         float64 `json:"price" binding:"required,gt=0"`
		Description   string  `json:"description"`
		StockQuantity int     `json:"stockQuantity" binding:"required,gte=0"`
		ReorderLevel  int     `json:"reorderLevel" binding:"required,gte=0,ltfield=StockQuantity"`
	}

	err := c.ShouldBindJSON(&updatedProductData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validate the input data

	// Check if the product exists
	var existingProduct models.Product
	err = initializer.DB.First(&existingProduct, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	// Check if the updated product name is unique
	isDuplicate := validators.IsProductNameDuplicate(updatedProductData.Name)
	if isDuplicate {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product with this name already exists",
		})
		return
	}

	// Check if stock quantity is greater than or equal to reorder level
	if updatedProductData.StockQuantity < updatedProductData.ReorderLevel {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Stock quantity must be greater than or equal to reorder level",
		})
		return
	}

	// Update the product in the database
	initializer.DB.Model(&existingProduct).Updates(models.Product{
		Name:          updatedProductData.Name,
		Category:      updatedProductData.Category,
		Price:         updatedProductData.Price,
		Description:   updatedProductData.Description,
		StockQuantity: updatedProductData.StockQuantity,
		ReorderLevel:  updatedProductData.ReorderLevel,
	})

	// Return a JSON response with the updated product
	c.JSON(http.StatusOK, gin.H{
		"updatedProduct": existingProduct,
	})
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

func DeleteProduct(c *gin.Context) {
	// Extract product ID from the request parameters
	id := c.Param("id")

	// Delete the product from the database
	err := initializer.DB.Delete(&models.Product{}, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete product",
		})
		return
	}

	// Return a JSON response indicating success
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
