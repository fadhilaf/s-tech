package delivery

import (
	"net/http"
	"net/url"

	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/google/uuid"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *orderHandler) CreateOrder(ctx *gin.Context) {
	var req model.CreateOrderFormRequest

	ok := utils.BindFormAndValidate(ctx, &req)
	if !ok {
		return
	}

	//convert productId dari string ke uuid
	productId, err := uuid.Parse(req.ProductID)
	if err != nil {
		return
	}

	// Ambil data user
	userId := utils.GetUserIdFromContext(ctx)
	if userId == uuid.Nil {
		return
	}

	res := handler.usecase.CreateOrder(model.CreateOrderRequest{
		UserID:      userId,
		ProductID:   productId,
		Quantity:    req.Quantity,
		Description: req.Description,
	})
	utils.SaveResponse(ctx, res.Message)

	// Gaya REST API
	// ctx.JSON(res.Status, res)

	// Gaya MVC

	var location url.URL
	location = url.URL{Path: "/pesan", RawQuery: "id=" + req.ProductID}

	if res.Status == http.StatusCreated {
		location = url.URL{Path: "/pesanan"}
	}

	ctx.Redirect(http.StatusFound, location.RequestURI())
}
