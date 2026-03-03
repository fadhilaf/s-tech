package usecase

import (
	"context"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
)

func (usecase *authUsecaseImpl) UserLogin(req model.LoginRequest) model.WebServiceResponse {
	user, err := usecase.Store.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		return utils.ToWebServiceResponse("Email belum terdaftar", http.StatusNotFound, nil)
	}

	if err := utils.ComparePassword(req.Password, user.PasswordHash); err != nil {
		return utils.ToWebServiceResponse("Password salah", http.StatusUnauthorized, nil)
	}

	//gaya REST API
	// return utils.ToWebServiceResponse("Login berhasil", http.StatusOK, gin.H{
	// 	"user": user,
	// })

	//gaya MVC
	modelUser := model.User{ID: user.ID, Name: user.Name, Email: user.Email, Address: user.Address, Phone: user.Phone}
	return utils.ToWebServiceResponse("Login berhasil", http.StatusOK, gin.H{
		"user": modelUser,
	})
}
