package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"runtime"
)

var router = gin.Default()

// Run will start the server
func Run() {
	getRoutes()
	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})
	router.LoadHTMLGlob("./public/html/*")

	router.Static("/public", "./public")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	port := os.Getenv("PORT")

	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	if os.Getenv("SSL") == "TRUE" {

		//Generated using sh generate-certificate.sh
		SSLKeys := &struct {
			CERT string
			KEY  string
		}{
			CERT: "./cert/myCA.cer",
			KEY:  "./cert/myCA.key",
		}

		router.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
		router.Run(":" + port)
	}
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	v1 := router.Group("/v1")
	addUserRoutes(v1)
}
