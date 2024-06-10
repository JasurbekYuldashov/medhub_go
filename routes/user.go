package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurbekyuldashov/medhub_go/controllers"
	"github.com/jmoiron/sqlx"
)

func addUserRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	ping := rg.Group("/users")

	controller1 := controllers.UserController{
		Database: db,
	}

	ping.GET("/", controller1.GetAll)
	ping.POST("/", controller1.CreateUser)
}
