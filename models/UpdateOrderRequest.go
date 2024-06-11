package models

type UpdateOrderRequest struct {
	BuyerName string           `json:"buyer_name" binding:"required"`
	Items     []OrderItemInput `json:"items"`
}

type OrderItemInput struct {
	ID        uint `json:"id"`
	ProductID uint `json:"product_id"`
	Count     int  `json:"count"`
	Delete    bool `json:"_delete"`
}
