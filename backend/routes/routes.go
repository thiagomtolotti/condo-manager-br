package routes

import (
	"backend/controllers"
	apartamentoController "backend/controllers/apartamento"
	moradorController "backend/controllers/morador"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", controllers.PingController)

	registerApartamentoRoutes(router)
	registerMoradorRoutes(router)
}

func registerApartamentoRoutes(router *gin.Engine) {
	router.GET("/apartamento", apartamentoController.Get)
	router.POST("/apartamento", apartamentoController.Create)
	router.DELETE("/apartamento/:id", apartamentoController.Delete)
}

func registerMoradorRoutes(router *gin.Engine) {
	router.POST("/morador", moradorController.Create)
}
