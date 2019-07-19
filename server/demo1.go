package server

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/henglory/Demo_Golang_v0.0.1/handler"
	"github.com/henglory/Demo_Golang_v0.0.1/service"
	"github.com/henglory/Demo_Golang_v0.0.1/spec"
)

func demo1(s service.Service, c *gin.Context) {
	var req spec.Demo1Req
	b, err := c.GetRawData()
	if err != nil {
		c.JSON(500, errorResponse{
			StatusCode: 99,
			StatusDesc: "MISS FORMAT",
			Request:    string(b),
		})
		return
	}
	err = json.Unmarshal(b, &req)
	if err != nil {
		c.JSON(500, errorResponse{
			StatusCode: 99,
			StatusDesc: "MISS FORMAT",
			Request:    string(b),
		})
		return
	}
	res := handler.Demo1(s, req)
	c.JSON(200, res)
}
