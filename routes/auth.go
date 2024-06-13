package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurbekyuldashov/medhub_go/controllers"
)

func addAuthRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/token")

	controller1 := new(controllers.AuthController)

	ping.POST("/refresh", controller1.Refresh)
}
