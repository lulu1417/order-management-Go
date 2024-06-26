package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"order-management-Go/database"
	"order-management-Go/models"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.NO = generateOrderNo()

	for i := range order.Items {
		productID := order.Items[i].ProductID
		if !(checkProductExist(productID)) {
			c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
			return
		}
		count := order.Items[i].Count
		var item models.OrderItem
		item.OrderID = order.ID
		item.ProductID = productID
		item.Count = count
		database.GetDB().Create(&item)
	}

	database.GetDB().Create(&order)
	c.JSON(http.StatusOK, gin.H{"success": "true"})
}

func generateOrderNo() string {
	return uuid.New().String()
}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	database.GetDB().Preload("Items.Product").Find(&orders)

	for i := range orders {
		orders[i].Total = orders[i].CalculateTotal()
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.GetDB().Preload("Items.Product").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	order.Total = order.CalculateTotal()
	c.JSON(http.StatusOK, order)
}

func UpdateOrder(c *gin.Context) {
	var order models.Order
	var requestData models.UpdateOrderRequest

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderID := c.Param("id")

	tx := database.GetDB().Begin()
	if err := tx.First(&order, orderID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := tx.Model(&order).Update("buyer_name", requestData.BuyerName).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, item := range requestData.Items {
		if item.ID != 0 {
			if !checkItemExist(item.ID, orderID) {
				tx.Rollback()
				c.JSON(http.StatusNotFound, gin.H{"error": "OrderItem not found"})
				return
			}

			if item.Delete {
				if err := tx.Delete(&models.OrderItem{}, item.ID).Error; err != nil {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
			} else {
				if item.ProductID != 0 && !checkProductExist(item.ProductID) {
					tx.Rollback()
					c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
					return
				}
				var orderItem models.OrderItem

				if err := tx.Model(&orderItem).Updates(item).Error; err != nil {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
			}
		} else {
			if !checkProductExist(item.ProductID) {
				tx.Rollback()
				c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
				return
			}

			if item.Count < 1 {
				tx.Rollback()
				c.JSON(http.StatusNotFound, gin.H{"message": "item count must greater than 0"})
				return
			}

			newItem := models.OrderItem{
				OrderID:   order.ID,
				ProductID: item.ProductID,
				Count:     item.Count,
			}

			if err := tx.Create(&newItem).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	database.GetDB().Preload("Items.Product").First(&order, orderID)
	order.Total = order.CalculateTotal()
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
	c.JSON(http.StatusOK, gin.H{"success": "true"})
}

func checkProductExist(id uint) bool {
	var product models.Product
	if err := database.GetDB().First(&product, id).Error; err != nil {
		return false
	}
	return true
}

func checkItemExist(itemId uint, orderId string) bool {
	var item models.OrderItem
	if err := database.GetDB().Where("id = ? AND order_id = ?", itemId, orderId).First(&item).Error; err != nil {
		return false
	}
	return true
}
