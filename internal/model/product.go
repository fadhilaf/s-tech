package model

import (
	"github.com/google/uuid"
)

type Product struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	CurrentPrice int32     `json:"current_price" db:"current_price"`
	Stock        int32     `json:"stock" db:"stock"`
	IsService    bool      `json:"is_service" db:"is_service"`
	Description  string    `json:"description" db:"description"`
	Image        string    `json:"image" db:"image_url"`
}

type CreateProductNoFileFormRequest struct {
	Name        string `form:"name" binding:"required"`
	Price       int32  `form:"price" binding:"required"`
	Stock       int32  `form:"stock" binding:"required"`
	SupplierID  string `form:"supplier_id" binding:"required"`
	IsService   bool   `form:"is_service" default:"false"`
	Description string `form:"description" binding:"required"`
}
type ProductNoFile struct {
	Name        string    `json:"name" binding:"required"`
	Price       int32     `json:"price" binding:"required"`
	Stock       int32     `json:"stock" binding:"required"`
	SupplierID  uuid.UUID `json:"supplier_id" binding:"required"`
	IsService   bool      `json:"is_service" default:"false"`
	Description string    `json:"description" binding:"required"`
}

type CreateProductRequest struct {
	ProductNoFile ProductNoFile
	Image         string
}

type GetProductByIdRequest struct {
	ID uuid.UUID
}

type GetProductByKeywordRequest struct {
	Keyword string `form:"search"`
}

type UpdateProductDetailsRequest struct {
	ID          uuid.UUID
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
	Image       string
}

type UpdateProductPriceRequest struct {
	ProductID    uuid.UUID
	Price        int32  `form:"price" binding:"required"`
	EffectiveDate string `form:"effective_date" binding:"required"`
}

type GetProductChronologyRequest struct {
	ProductID uuid.UUID
}

type GetProductPricesRequest struct {
	ProductID uuid.UUID
}
