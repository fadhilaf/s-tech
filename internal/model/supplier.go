package model

import (
	"github.com/google/uuid"
)

type Supplier struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateSupplierRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}
