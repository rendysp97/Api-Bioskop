package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rendysp97/Api-Bioskop/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskop", controllers.CreateDataBioskop)

	router.GET("/bioskop", controllers.GetAllData)

	router.GET("/bioskop/:id", controllers.GetDetailBioskop)

	router.PUT("/bioskop/:id", controllers.UpdateDataBioskop)

	router.DELETE("/bioskop/:id", controllers.DeleteDataBioskop)

	return router
}
