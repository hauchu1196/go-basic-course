// internal/api/routes.go
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hauchu1196/ecombase/internal/api/handlers"
	"github.com/hauchu1196/ecombase/internal/repository"
	"github.com/hauchu1196/ecombase/internal/service"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Order routes
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("", orderHandler.CreateOrder)
		orderRoutes.GET("/:id", orderHandler.GetOrder)
		orderRoutes.GET("", orderHandler.ListOrders)
		orderRoutes.PUT("/:id", orderHandler.UpdateOrder)
		orderRoutes.DELETE("/:id", orderHandler.DeleteOrder)
	}

	// Add other entity routes here
	// Example:
	// productRepo := repository.NewProductRepository(db)
	// productService := service.NewProductService(productRepo)
	// productHandler := handlers.NewProductHandler(productService)
	//
	// productRoutes := router.Group("/products")
	// {
	//     productRoutes.POST("", productHandler.CreateProduct)
	//     ...
	// }
}
