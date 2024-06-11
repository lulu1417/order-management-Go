package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-management-Go/database"
	"order-management-Go/models"
	"strconv"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.GetDB().Create(&order)
	c.JSON(http.StatusOK, order)
}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	database.GetDB().Preload("OrderItems.Product").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.GetDB().Preload("OrderItems.Product").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	uintID, _ := strconv.ParseUint(id, 10, 32)
	var order models.Order
	if err := database.GetDB().First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.ID = uint(uintID)
	database.GetDB().Save(&order)
	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.GetDB().First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}

	database.GetDB().Delete(&order)
	c.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}
