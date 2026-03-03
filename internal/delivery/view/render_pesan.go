package delivery

import (
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (handler *viewHandler) RenderPesan(c *gin.Context) {
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

	// Ambil data product
	productIdQuery := c.Query("id")
	productId, err := uuid.Parse(productIdQuery)
	if err != nil {
		c.HTML(http.StatusBadRequest, "400.gohtml", gin.H{"Message": "Tidak ada produk yang dipilih"})
		c.Abort()
		return
	}

	req := model.GetProductByIdRequest{ID: productId}
	res := handler.usecase.GetProductById(req)

	if res.Status == http.StatusNotFound {
		c.HTML(http.StatusNotFound, "404.gohtml", gin.H{})
		c.Abort()
		return
	}

	product, ok := res.Data["product"].(model.Product)
	if !ok {
		c.HTML(http.StatusInternalServerError, "500.gohtml", gin.H{})
		c.Abort()
	}

	c.HTML(http.StatusOK, "pesan.gohtml", gin.H{
		"Message": message,
		"Name":    name,
		"Product": product,
	})
}
