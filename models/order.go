package models

import "time"

type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	NO        string      `json:"no"`
	BuyerName string      `json:"buyer_name" binding:"required"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Total     float64     `json:"total_amount" gorm:"-"`
	Items     []OrderItem `json:"items"`
}

func (o *Order) CalculateTotal() float64 {
	var total float64
	for _, item := range o.Items {
		total += item.Product.Price * float64(item.Count)
	}
	return total
}
