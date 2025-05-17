package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", controllers.PingController)

	registerApartamentoRoutes(router)
}

func registerApartamentoRoutes(router *gin.Engine) {
	router.POST("/apartamento", controllers.CreateApartamento)
}
