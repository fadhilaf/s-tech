package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID `json:"id" db:"id"`
	Name    string    `json:"name" db:"name"`
	Email   string    `json:"email" db:"email"`
	Address string    `json:"address" db:"address"`
	Phone   string    `json:"phone" db:"phone"`
}

type CreateUserRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Address  string `json:"address" form:"address" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required"`
}

type GetUserByIdRequest struct {
	ID uuid.UUID `json:"id" query:"id" binding:"required,uuid"`
}

type UpdateUserRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	Phone   string `json:"phone" form:"phone" binding:"required"`
}

type DeleteUserRequest struct {
	ID uuid.UUID `json:"id" form:"id" binding:"required,uuid"`
}
