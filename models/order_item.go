package models

import "time"

type OrderItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderID   uint      `json:"order_id" binding:"required"`
	ProductID uint      `json:"product_id" binding:"required"`
	Count     int       `json:"count" binding:"required, min=0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
}
