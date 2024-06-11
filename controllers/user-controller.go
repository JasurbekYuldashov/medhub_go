package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurbekyuldashov/medhub_go/models"
	"log"
	"net/http"
)

type UserController struct {
}

var userModel = models.UserModel{}

func (u UserController) GetAll(c *gin.Context) {
	users, err := userModel.GetAll()
	if err != nil {
		log.Printf("error %d", err)
		SendResponse(c, http.StatusBadRequest, "error", nil, err)
		return
	}
	SendResponse(c, http.StatusOK, "success", users, nil)
}

func (u UserController) GetOne(c *gin.Context) {
	users, err := userModel.GetAll()
	if err != nil {
		log.Printf("error %d", err)
		SendResponse(c, http.StatusBadRequest, "error", nil, err)
		return
	}
	SendResponse(c, http.StatusOK, "success", users, nil)
}
