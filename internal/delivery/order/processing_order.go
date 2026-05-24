package delivery

import (
	"net/http"

	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/google/uuid"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *orderHandler) ProcessingOrder(ctx *gin.Context) {
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

	req := model.UpdateOrderStatusProcessingRequest{
		ID: orderId,
	}
	res := handler.usecase.ProcessingOrder(req)
	if res.Status == http.StatusOK {
		ctx.Header("HX-Refresh", "true")
	}

	ctx.JSON(res.Status, res)
}
