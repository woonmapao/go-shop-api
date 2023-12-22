package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/woonmapao/user-management/controllers"
	"github.com/woonmapao/user-management/services"
)

// SetupRouter configures and returns the router with all API routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	r.POST("/users", controllers.AddUser)
	r.GET("/users/:id", controllers.GetUserByID)
	r.GET("/users", controllers.GetAllUsers)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Order routes
	r.GET("/orders", controllers.GetAllOrders)
	r.GET("/orders/:id", controllers.GetOrderByID)
	r.POST("/orders", controllers.CreateOrder)
	r.PUT("/orders/:id", controllers.UpdateOrder)
	r.DELETE("/orders/:id", controllers.DeleteOrder)
	r.GET("/users/:id/orders", controllers.GetUserOrders)

	// Product routes
	r.GET("/products", controllers.GetAllProducts)
	r.GET("/products/:id", controllers.GetProductByID)
	r.POST("/products", controllers.CreateProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	// Search products
	r.GET("/search/products", services.SearchProducts)

	// Order Detail routes
	r.GET("/order-details", controllers.GetAllOrderDetails)
	r.GET("/order-details/:id", controllers.GetOrderDetailByID)
	r.POST("/order-details", controllers.CreateOrderDetail)
	r.PUT("/order-details/:id", controllers.UpdateOrderDetail)
	r.DELETE("/order-details/:id", controllers.DeleteOrderDetail)
	r.GET("/order-details/by-order/:orderId", controllers.GetOrderDetailsByOrderID)

	// Purchase Product
	r.POST("/purchase", services.PurchaseProduct)

	return r
}
