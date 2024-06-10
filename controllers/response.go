package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code  int         `json:"statusCode"`
	Msg   interface{} `json:"message,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

func ResponseError(c *gin.Context, statusCode int, error interface{}, message string) {
	c.JSON(http.StatusOK, &ResponseData{
		Code:  statusCode,
		Msg:   message,
		Error: error,
	})
}

func ResponseSuccess(c *gin.Context, statusCode int, data interface{}, message string) {
	message = ""
	data = nil
	c.JSON(statusCode, &ResponseData{
		Code: statusCode,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

func SendResponse(c *gin.Context, status int, message string, data interface{}, err interface{}) {
	response := ResponseData{
		Code:  status,
		Msg:   message,
		Data:  data,
		Error: err,
	}
	c.JSON(status, response)
}

func ResponseErrorWithMsg(c *gin.Context, statusCode int, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: statusCode,
		Msg:  msg,
		Data: nil,
	})
}
