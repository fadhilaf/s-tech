package usecase

import (
	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/repository"
	"github.com/google/uuid"
)

type UserUsecase interface {
	CreateUser(model.CreateUserRequest) model.WebServiceResponse
	// DeleteUser(model.DeleteUserRequest) model.WebServiceResponse
	GetProfile(userId uuid.UUID) model.WebServiceResponse
	// ListUser() model.WebServiceResponse
	// UpdateUser(model.UpdateUserRequest) model.WebServiceResponse
}

var _ UserUsecase = &userUsecaseImpl{}

func NewUserUsecase(store repository.Store) UserUsecase {
	return &userUsecaseImpl{
		Store: store,
	}
}

type userUsecaseImpl struct {
	repository.Store
}
