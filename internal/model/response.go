package model

import "github.com/gin-gonic/gin"

type WebServiceResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    gin.H  `json:"data"`
}

type ValidationErrorWebServiceResponse struct {
	WebServiceResponse
	DetailErrors []string `json:"errors"`
}
