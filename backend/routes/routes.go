package routes

import (
	"backend/controllers"
	apartamentoController "backend/controllers/apartamento"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", controllers.PingController)

	registerApartamentoRoutes(router)
}

func registerApartamentoRoutes(router *gin.Engine) {
	router.GET("/apartamento", apartamentoController.Get)
	router.POST("/apartamento", apartamentoController.Create)
}
