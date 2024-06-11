package models

import "time"

type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	BuyerName  string      `json:"buyer_name"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Total      float64     `json:"total_amount" gorm:"-"`
	OrderItems []OrderItem `json:"items"`
}

func (o *Order) CalculateTotal() {
	var total float64
	for _, item := range o.OrderItems {
		total += item.Product.Price * float64(item.Count)
	}
	o.Total = total
}
