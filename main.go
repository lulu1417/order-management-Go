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

	router.POST("/products", controllers.CreateProduct)
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProduct)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrders)
	router.GET("/orders/:id", controllers.GetOrder)
	router.PUT("/orders/:id", controllers.UpdateOrder)
	router.DELETE("/orders/:id", controllers.DeleteOrder)

	router.Run(":8000")
}
