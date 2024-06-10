package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurbekyuldashov/medhub_go/models"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

type UserController struct {
	Database *sqlx.DB
}

func (u UserController) GetAll(c *gin.Context) {
	var users []models.User = make([]models.User, 0)
	err := u.Database.Select(&users, "SELECT id, full_name, email FROM users")
	if err != nil {
		SendResponse(c, http.StatusBadRequest, "error", nil, err.Error())
		return
	}

	SendResponse(c, http.StatusOK, "success", users, nil)
}

func (u UserController) CreateUser(c *gin.Context) {
	var user1 models.User
	if err := c.BindJSON(&user1); err != nil {
		SendResponse(c, http.StatusBadRequest, "error", nil, err.Error())
	}

	user, err := u.Database.Queryx("INSERT INTO users (email, password, full_name, updated_at) VALUES ($1, $2, $3, $4);", user1.Email, user1.Password, user1.FullName, time.Now())
	if err != nil {
		SendResponse(c, http.StatusBadRequest, "error", nil, err.Error())
		return
	}
	log.Println(user.Scan())
	SendResponse(c, http.StatusOK, "success", user, nil)
}
