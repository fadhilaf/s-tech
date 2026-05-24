package delivery

import (
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
)

func (delivery *authHandler) GetProfile(ctx *gin.Context) {
	userId := utils.GetUserIdFromContext(ctx)
	isAdmin := utils.GetAdminFromContext(ctx)

	// We rely on ShouldAuth middleware to ensure at least one is valid
	res := delivery.usecase.GetProfile(userId, isAdmin)
	ctx.JSON(res.Status, res)
}
