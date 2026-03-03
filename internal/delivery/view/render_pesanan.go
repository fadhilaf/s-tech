package delivery

import (
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (handler *viewHandler) RenderPesanan(c *gin.Context) {
	message := utils.GetResponse(c)

	// Ambil data user
	userId := utils.GetUserIdFromContext(c)
	if userId == uuid.Nil {
		return
	}

	resUser := handler.usecase.GetUserById(model.GetUserByIdRequest{ID: userId})
	if resUser.Status != http.StatusOK {
		return
	}
	user, ok := resUser.Data["user"].(model.User)

	var name string
	if ok {
		name = user.Name
	}

	// Ambil data pesanan
	resOrder := handler.usecase.GetOrderByUserId(model.GetOrderByUserIdRequest{UserID: userId})
	var orders []model.Order
	if resOrder.Status == http.StatusOK {
		orders, _ = resOrder.Data["orders"].([]model.Order)
	}

	adminPhone := handler.AdminPhone

	c.HTML(http.StatusOK, "pesanan.gohtml", gin.H{
		"Message":    message,
		"Name":       name,
		"AdminPhone": adminPhone,
		"Orders":     orders,
	})
}
