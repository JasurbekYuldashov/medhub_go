package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurbekyuldashov/medhub_go/forms"
	"github.com/jasurbekyuldashov/medhub_go/models"
	"log"
	"net/http"
)

type UserController struct {
}

var userModel = new(models.UserModel)
var userForm = new(forms.UserForm)

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

// getUserID ...
func getUserID(c *gin.Context) (userID int64) {
	//MustGet returns the value for the given key if it exists, otherwise it panics.
	return c.MustGet("userID").(int64)
}

// Login ...
func (ctrl UserController) Login(c *gin.Context) {
	var loginForm forms.LoginForm

	if validationErr := c.ShouldBindJSON(&loginForm); validationErr != nil {
		message := userForm.Login(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, token, err := userModel.Login(loginForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid login details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "user": user, "token": token})
}

// Register ...
func (ctrl UserController) Register(c *gin.Context) {
	var registerForm forms.RegisterForm

	if validationErr := c.ShouldBindJSON(&registerForm); validationErr != nil {
		message := userForm.Register(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, err := userModel.Register(registerForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered", "user": user})
}

// Logout ...
func (ctrl UserController) Logout(c *gin.Context) {

	au, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User not logged in"})
		return
	}

	deleted, delErr := authModel.DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
