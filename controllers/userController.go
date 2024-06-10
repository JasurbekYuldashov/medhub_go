package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasurbekyuldashov/medhub_go/db"
	"github.com/jasurbekyuldashov/medhub_go/models"
	"log"
	"net/http"
	"time"
)

type UserController struct {
}

func (u UserController) GetAll(c *gin.Context) {
	d := db.GetDB()
	var users []models.User = make([]models.User, 0)
	err := d.Select(&users, "SELECT id, full_name, email FROM users")
	if err != nil {
		SendResponse(c, http.StatusBadRequest, "error", nil, err.Error())
		return
	}

	SendResponse(c, http.StatusOK, "success", users, nil)
}

func (u UserController) CreateUser(c *gin.Context) {
	d := db.GetDB()
	var user1 models.User
	if err := c.BindJSON(&user1); err != nil {
		SendResponse(c, http.StatusBadRequest, "error", nil, err.Error())
	}

	user, err := d.Queryx("INSERT INTO users (email, password, full_name, updated_at) VALUES ($1, $2, $3, $4);", user1.Email, user1.Password, user1.FullName, time.Now())
	if err != nil {
		SendResponse(c, http.StatusBadRequest, "error", nil, err.Error())
		return
	}

	for user.Next() {
		var user1 models.UserGet
		if err := user.StructScan(&user1); err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("User: %+v\n", user1)
	}
	SendResponse(c, http.StatusOK, "success", user1, nil)
}
