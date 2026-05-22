package utils

import (
	"fmt"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/gin-gonic/gin"
)

// REST API
func ToWebServiceResponse(message string, status int, data gin.H) model.WebServiceResponse {
	return model.WebServiceResponse{
		Message: message,
		Status:  status,
		Data:    data,
	}
}

func ToDetailedErrorWebServiceResponse(message string, status int, errors map[string]string) model.DetailedErrorWebServiceResponse {
	return model.DetailedErrorWebServiceResponse{
		WebServiceResponse: model.WebServiceResponse{
			Message: message,
			Status:  status,
			Data:    nil,
		},
		DetailErrors: errors,
	}
}

// MVC
func SaveResponse(c *gin.Context, message string) {
	c.SetCookie("response", message, 60, "/", "localhost", false, false)
}

func GetResponse(c *gin.Context) string {
	cookie, err := c.Cookie("response")

	fmt.Println(cookie)

	if err != nil {
		return ""
	}

	c.SetCookie("response", "", -1, "/", "localhost", false, true)
	return cookie
}
