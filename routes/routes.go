package routes

import (
	"github.com/kotopanjang/aqilliz_assesment/controllers"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	router.GET("/register", controllers.Register)
	router.GET("/retrieve", controllers.Retrieve)
	router.GET("/redirect", controllers.Redirect)

	return router
}
