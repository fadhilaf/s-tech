package delivery

import (
	usecase "github.com/fadhilaf/s-tech/internal/usecase/auth"
	"github.com/gin-gonic/gin"
)

type AuthDelivery interface {
	UserLogin(ctx *gin.Context)
	AdminLogin(ctx *gin.Context)
	Logout(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
}

var _ AuthDelivery = &authHandler{}

func NewAuthDelivery(usecase usecase.AuthUsecase) AuthDelivery {
	return &authHandler{
		usecase: usecase,
	}
}

type authHandler struct {
	usecase usecase.AuthUsecase
}
