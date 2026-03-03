package model

import (
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Price       string    `json:"price" db:"price"`
	Stock       int32     `json:"stock" db:"stock"`
	IsService   bool      `json:"is_service" db:"is_service"`
	Description string    `json:"description" db:"description"`
	Image       string    `json:"image" db:"image_url"`
}

type CreateProductNoFileRequest struct {
	Name        string `form:"name" binding:"required"`
	Price       int32  `form:"price" binding:"required"`
	Stock       int32  `form:"stock" binding:"required"`
	IsService   bool   `form:"is_service" default:"false"`
	Description string `form:"description" binding:"required"`
}

type CreateProductRequest struct {
	NotFile CreateProductNoFileRequest
	Image   string
}

type GetProductByIdRequest struct {
	ID uuid.UUID
}

type GetProductByKeywordRequest struct {
	Keyword string `form:"search"`
}
