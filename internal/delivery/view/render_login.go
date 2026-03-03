package delivery

import (
	"net/http"

	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderLogin(c *gin.Context) {
	message := utils.GetResponse(c)

	c.HTML(http.StatusOK, "login.gohtml", gin.H{
		"Message": message,
	})
}
