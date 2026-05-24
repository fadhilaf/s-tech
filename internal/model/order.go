package model

import (
	"github.com/google/uuid"
)

const (
	OrderStatusPending    string = "pending"
	OrderStatusProcessing string = "processing"
	OrderStatusDelivered  string = "delivered"
)

type Order struct {
	ID             uuid.UUID `json:"id"`
	ProductPriceID uuid.UUID `json:"product_price_id"`
	ProductID      uuid.UUID `json:"product_id"`
	ProductName    string    `json:"product_name"`
	ProductPrice   int32     `json:"product_price"`
	IsService      bool      `json:"is_service"`
	BuyerID        uuid.UUID `json:"buyer_id"`
	BuyerName      string    `json:"buyer_name"`
	BuyerAddress   string    `json:"buyer_address"`
	BuyerPhone     string    `json:"buyer_phone"`
	Quantity       int32     `json:"quantity"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
}

type CreateOrderFormRequest struct {
	// di struct untuk bind request, dak biso pake data type uuid.UUID, harus pake string. validasi uuid nyo tetep ado di bagian binding itu
	ProductID   string `json:"product_id" form:"product_id" binding:"required,uuid"`
	Quantity    int32  `json:"quantity" form:"quantity" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}

type CreateOrderRequest struct {
	UserID      uuid.UUID
	ProductID   uuid.UUID
	Quantity    int32
	Description string
}

type GetOrderByUserIdRequest struct {
	UserID uuid.UUID
}

type UpdateOrderStatusFormRequest struct {
	ID string `json:"id" form:"id" binding:"required,uuid"`
}

type UpdateOrderStatusProcessingRequest struct {
	ID uuid.UUID
}

type UpdateOrderStatusDeliveredRequest struct {
	ID     uuid.UUID
	UserID uuid.UUID
}
