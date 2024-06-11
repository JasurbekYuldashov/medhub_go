package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurbekyuldashov/medhub_go/controllers"
)

func addUserRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/users")

	controller1 := controllers.UserController{}

	ping.GET("/", controller1.GetAll)
	ping.GET("/:id", controller1.GetOne)
	//ping.POST("/", controller1.CreateUser)
}
