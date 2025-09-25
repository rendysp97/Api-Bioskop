package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rendysp97/Api-Bioskop/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskop", controllers.CreateDataBioskop)

	return router
}
