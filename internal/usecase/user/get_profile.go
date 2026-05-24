package usecase

import (
	"context"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (usecase *userUsecaseImpl) GetProfile(userId uuid.UUID) model.WebServiceResponse {
	userDb, err := usecase.Store.GetUserById(context.Background(), userId)
	if err != nil {
		return utils.ToWebServiceResponse("User tidak ditemukan", http.StatusNotFound, nil)
	}

	user := model.User{
		ID:      userDb.ID,
		Name:    userDb.Name,
		Email:   userDb.Email,
		Address: userDb.Address,
		Phone:   userDb.Phone,
	}

	return utils.ToWebServiceResponse("Berhasil mendapatkan profile", http.StatusOK, gin.H{
		"user": user,
	})
}
