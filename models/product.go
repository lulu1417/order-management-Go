package models

import (
	"time"
)

type Product struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name" binding:"required"`
	Price     float64   `json:"price" binding:"required,min=0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
