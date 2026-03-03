package usecase

import (
	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/repository"
)

type AuthUsecase interface {
	UserLogin(model.LoginRequest) model.WebServiceResponse
	AdminLogin(model.LoginRequest) model.WebServiceResponse
}

var _ AuthUsecase = &authUsecaseImpl{}

func NewAuthUsecase(store repository.Store, adminEmail string, adminPassword string) AuthUsecase {
	return &authUsecaseImpl{
		Store: store,
		AdminEmail: adminEmail,
		AdminPassword: adminPassword,
	}
}

type authUsecaseImpl struct {
	repository.Store
	AdminEmail string
	AdminPassword string
}
