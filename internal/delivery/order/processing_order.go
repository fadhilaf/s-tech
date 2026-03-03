package delivery

import (
	"net/http"
	"net/url"

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
	utils.SaveResponse(ctx, res.Message)

	// Gaya REST API
	// ctx.JSON(res.Status, res)

	// Gaya MVC
	location := url.URL{Path: "/admin/pesanan"}
	ctx.Redirect(http.StatusFound, location.RequestURI())
}
