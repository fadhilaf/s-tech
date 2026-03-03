package usecase

import (
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
)

func (usecase *authUsecaseImpl) AdminLogin(req model.LoginRequest) model.WebServiceResponse {

	adminEmail := usecase.AdminEmail
	adminPassword := usecase.AdminPassword

	if req.Email != adminEmail {
		return utils.ToWebServiceResponse("Email salah", http.StatusUnauthorized, nil)
	}

	if req.Password != adminPassword {
		return utils.ToWebServiceResponse("Password salah", http.StatusUnauthorized, nil)
	}

	return utils.ToWebServiceResponse("Berhasil masuk sebagai admin", http.StatusOK, nil)
}
