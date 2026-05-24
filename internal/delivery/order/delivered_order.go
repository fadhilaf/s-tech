package delivery

import (
	"net/http"

	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/google/uuid"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *orderHandler) DeliveredOrder(ctx *gin.Context) {
	var reqForm model.UpdateOrderStatusFormRequest

	ok := utils.BindFormAndValidate(ctx, &reqForm)
	if !ok {
		return
	}

	//convert orderId dari string ke uuid
	orderId, err := uuid.Parse(reqForm.ID)
	if err != nil {
		return
	}

	// Ambil data user
	userId := utils.GetUserIdFromContext(ctx)
	if userId == uuid.Nil {
		return
	}

	req := model.UpdateOrderStatusDeliveredRequest{
		ID:     orderId,
		UserID: userId,
	}
	res := handler.usecase.DeliveredOrder(req)
	if res.Status == http.StatusOK {
		ctx.Header("HX-Refresh", "true")
	}

	ctx.JSON(res.Status, res)
}
