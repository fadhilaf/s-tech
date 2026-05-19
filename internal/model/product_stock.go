package model

import (
	"github.com/google/uuid"
)

type ProductStock struct {
	ID         uuid.UUID `json:"id"`
	ProductID  uuid.UUID `json:"product_id"`
	SupplierID uuid.UUID `json:"supplier_id"`
	IsAdd      bool      `json:"is_add"`
	Quantity   int32     `json:"quantity"`
	Price      int32     `json:"price"`
}

type ReceiveProductStockRequest struct {
	ProductID  string `json:"product_id" form:"product_id" binding:"required,uuid"`
	SupplierID string `json:"supplier_id" form:"supplier_id" binding:"required,uuid"`
	Quantity   int32  `json:"quantity" form:"quantity" binding:"required"`
	Price      int32  `json:"price" form:"price" binding:"required"`
}
