package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/models"
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
	isDuplicate := isProductNameDuplicate(body.Name)
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

// Check if a product with the given name already exists in the database
func isProductNameDuplicate(name string) bool {
	var existingProduct models.Product
	err := initializer.DB.Where("name = ?", name).First(&existingProduct).Error
	return err == nil
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
	isDuplicate := isProductNameDuplicate(updatedProductData.Name)
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

func PurchaseProduct(c *gin.Context) {
	// Extract user and product information from the request
	var purchaseData struct {
		UserID    int `json:"userId" binding:"required"`
		ProductID int `json:"productId" binding:"required"`
		Quantity  int `json:"quantity" binding:"required,gte=1"`
	}

	err := c.ShouldBindJSON(&purchaseData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	purchase := struct {
		UserID    int
		ProductID int
		Quantity  int
	}{
		UserID:    purchaseData.UserID,
		ProductID: purchaseData.ProductID,
		Quantity:  purchaseData.Quantity,
	}

	// Validate input data
	err = validatePurchaseData(purchase)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Handle the purchase (update stock, create order)
	err = updateStockAndCreateOrder(purchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process the purchase",
		})
		return
	}

	// Return a JSON response indicating the success of the purchase
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

// Validate purchase data
func validatePurchaseData(data struct{ UserID, ProductID, Quantity int }) error {

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

// Check if a user with the given ID exists in the database
func userExists(userID int) bool {
	var user models.User
	result := initializer.DB.First(&user, userID)
	return result.RowsAffected > 0
}

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

// Update stock and create an order in the database
func updateStockAndCreateOrder(data struct{ UserID, ProductID, Quantity int }) error {

	// Update stock
	err := updateStock(data.ProductID, data.Quantity)
	if err != nil {
		return err
	}

	// Create order detail and order records
	err = createOrderAndDetail(data)
	if err != nil {
		return err
	}

	return nil
}

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

// Create order and order detail records in the database
func createOrderAndDetail(data struct{ UserID, ProductID, Quantity int }) error {

	subtotal := calculateSubtotal(data.ProductID, data.Quantity)

	// Create order
	order := models.Order{
		UserID:      data.UserID,
		OrderDate:   time.Now(),
		TotalAmount: subtotal,
		Status:      "Pending", // or any initial status
	}
	initializer.DB.Create(&order)

	// Create order detail
	orderDetail := models.OrderDetail{
		OrderID:   int(order.ID),
		ProductID: data.ProductID,
		Quantity:  data.Quantity,
		Subtotal:  subtotal,
	}
	initializer.DB.Create(&orderDetail)

	return nil
}

func calculateSubtotal(productID, quantity int) float64 {
	var product models.Product
	result := initializer.DB.First(&product, productID)
	if result.Error != nil {
		return 0
	}

	return float64(quantity) * product.Price
}
