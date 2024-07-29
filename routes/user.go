package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurbekyuldashov/medhub_go/controllers"
	"github.com/jasurbekyuldashov/medhub_go/db"
)

func addUserRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/user")

	controller1 := controllers.UserController{}
	var (
		db1                                               = db.GetDB()
		userTestController controllers.UserTestController = controllers.NewUserTestController(db1)
	)

	ping.GET("/index", userTestController.Index)
	ping.GET("/", controller1.GetAll)
	ping.GET("/:id", controller1.GetOne)
	ping.POST("/login", controller1.Login)
	ping.POST("/register", controller1.Register)
	ping.GET("/logout", controller1.Logout)
}
