package routes

import (
	"github.com/amarsinghrathour/oolio_backend_challenge/config"
	"github.com/amarsinghrathour/oolio_backend_challenge/controllers"
	"github.com/amarsinghrathour/oolio_backend_challenge/database"
	"github.com/amarsinghrathour/oolio_backend_challenge/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *database.DBInstance, cfg *config.Config) {
	// Service layer
	productService := services.NewProductService(db.DB)
	orderService := services.NewOrderService(productService)

	// Controller layer
	productController := controllers.NewProductController(productService)
	orderController := controllers.NewOrderController(orderService)

	// Routes
	productGroup := router.Group("/product")
	{
		productGroup.GET("", productController.GetAllProducts)
		productGroup.GET("/:productId", productController.GetProductByID)
	}
	orderGroup := router.Group("/order")
	{
		orderGroup.POST("", orderController.PlaceOrder)
	}
}
