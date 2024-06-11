package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthController struct {
}

func (u AuthController) GetAll(c *gin.Context) {
	users, err := userModel.GetAll()
	if err != nil {
		log.Printf("error %d", err)
		SendResponse(c, http.StatusBadRequest, "error", nil, err)
		return
	}
	SendResponse(c, http.StatusOK, "success", users, nil)
}

func (u AuthController) GetOne(c *gin.Context) {
	users, err := userModel.GetAll()
	if err != nil {
		log.Printf("error %d", err)
		SendResponse(c, http.StatusBadRequest, "error", nil, err)
		return
	}
	SendResponse(c, http.StatusOK, "success", users, nil)
}
