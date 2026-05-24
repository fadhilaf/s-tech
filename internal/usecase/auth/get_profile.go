package usecase

import (
	"context"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (usecase *authUsecaseImpl) GetProfile(userId uuid.UUID, isAdmin bool) model.WebServiceResponse {
	if isAdmin {
		adminUser := model.User{
			Name:    "Admin S-TECH",
			Email:   usecase.AdminEmail,
			IsAdmin: true,
		}
		return utils.ToWebServiceResponse("Berhasil mendapatkan profile admin", http.StatusOK, gin.H{
			"user": adminUser,
		})
	}

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
		IsAdmin: false,
	}

	return utils.ToWebServiceResponse("Berhasil mendapatkan profile user", http.StatusOK, gin.H{
		"user": user,
	})
}
