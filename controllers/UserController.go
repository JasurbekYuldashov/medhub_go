package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UserTestController interface {
	Index(ctx *gin.Context)
}
type userTestController struct {
	db *sqlx.DB
}

func NewUserTestController(dd *sqlx.DB) UserTestController {
	return &userTestController{
		db: dd,
	}
}

func (c userTestController) Index(ctx *gin.Context) {

}

func (c userTestController) indexAA(ctx *gin.Context) {

}
