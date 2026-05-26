package usecase

import (
	"context"

	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	repositoryModel "github.com/fadhilaf/s-tech/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	utils "github.com/fadhilaf/s-tech/internal/utils"
)

func (usecase *userUsecaseImpl) CreateUser(req model.CreateUserRequest) model.WebServiceResponse {
	_, err := usecase.Store.GetUserByEmail(context.Background(), req.Email)
	if err == nil {
		return utils.ToWebServiceResponse("Email sudah terdaftar", http.StatusConflict, gin.H{
			"errors": map[string]string{"email": "Email sudah terdaftar"},
		})
	}

	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return utils.ToWebServiceResponse("Fungsi hash password gagal", http.StatusInternalServerError, nil)
	}

	user, err := usecase.Store.CreateUser(context.Background(), repositoryModel.CreateUserParams{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
		Address:      req.Address,
		Phone:        req.Phone,
	})

	if err != nil {
		return utils.ToWebServiceResponse("Menambah user ke db gagal", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("User berhasil dibuat", http.StatusCreated, gin.H{"user": user})
}
