package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type WebServiceResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    gin.H  `json:"data"`
}

func main() {
	res := WebServiceResponse{
		Message: "Test",
		Status:  200,
		Data:    nil,
	}
	b, _ := json.Marshal(res)
	fmt.Println(string(b))

	res2 := WebServiceResponse{
		Message: "Test2",
		Status:  200,
		Data:    gin.H{"prices": nil},
	}
	b2, _ := json.Marshal(res2)
	fmt.Println(string(b2))
}
