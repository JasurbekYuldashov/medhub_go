package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurbekyuldashov/medhub_go/controllers"
)

func addUserRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/user")

	controller1 := controllers.UserController{}

	ping.GET("/", controller1.GetAll)
	ping.GET("/:id", controller1.GetOne)
	ping.POST("/login", controller1.Login)
	ping.POST("/register", controller1.Register)
	ping.GET("/logout", controller1.Logout)
}
