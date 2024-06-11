package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-management-Go/database"
	"order-management-Go/models"
	"strconv"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !checkNameUnique(product.Name) {
		c.JSON(http.StatusNotFound, gin.H{"message": "product name existed"})
		return
	}

	database.GetDB().Create(&product)
	c.JSON(http.StatusOK, gin.H{"success": "true"})
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	database.GetDB().Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.GetDB().First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	uintID, _ := strconv.ParseUint(id, 10, 32)
	var product models.Product
	if err := database.GetDB().First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = uint(uintID)
	database.GetDB().Save(&product)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.GetDB().First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	if !canDelete(product.ID) {
		c.JSON(http.StatusNotFound, gin.H{"message": "product exists in order items"})
		return
	}

	database.GetDB().Delete(&product)
	c.JSON(http.StatusOK, gin.H{"success": "true"})
}

func checkNameUnique(name string) bool {
	var count int
	var product models.Product
	database.GetDB().Where("name = ?", name).First(&product).Count(&count)
	return count == 0
}

func canDelete(productId uint) bool {
	var count int
	var item models.OrderItem
	database.GetDB().Where("product_id = ?", productId).First(&item).Count(&count)
	return count == 0
}
