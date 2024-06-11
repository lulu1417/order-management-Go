package main

import (
	"github.com/gin-gonic/gin"
	"order-management-Go/controllers"
	"order-management-Go/database"
	"order-management-Go/models"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	db.AutoMigrate(&models.Product{}, &models.Order{}, &models.OrderItem{})

	router := gin.Default()

	router.POST("/api/products", controllers.CreateProduct)
	router.GET("/api/products", controllers.GetProducts)
	router.GET("/api/products/:id", controllers.GetProduct)
	router.PUT("/api/products/:id", controllers.UpdateProduct)
	router.DELETE("/api/products/:id", controllers.DeleteProduct)
	router.POST("/api/orders", controllers.CreateOrder)
	router.GET("/api/orders", controllers.GetOrders)
	router.GET("/api/orders/:id", controllers.GetOrder)
	router.PUT("/api/orders/:id", controllers.UpdateOrder)
	router.DELETE("/api/orders/:id", controllers.DeleteOrder)

	router.Run(":8000")
}
