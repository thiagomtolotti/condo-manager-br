package routes

import (
	"backend/controllers"
	apartamentoController "backend/controllers/apartamento"
	moradorController "backend/controllers/morador"
	vagaController "backend/controllers/vaga"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", controllers.PingController)

	registerApartamentoRoutes(router)
	registerMoradorRoutes(router)
	registerVagaRoutes(router)
}

func registerApartamentoRoutes(router *gin.Engine) {
	router.GET("/apartamento", apartamentoController.Get)
	router.POST("/apartamento", apartamentoController.Create)
	router.DELETE("/apartamento/:id", apartamentoController.Delete)
}

func registerMoradorRoutes(router *gin.Engine) {
	router.GET("/morador", moradorController.Get)
	router.POST("/morador", moradorController.Create)
	router.PATCH("/morador/:cpf", moradorController.Patch)
	router.DELETE("/morador/:cpf", moradorController.Delete)
}

func registerVagaRoutes(router *gin.Engine) {
	router.GET("/vaga", vagaController.Get)
	router.POST("/vaga/:apartamento_id", vagaController.Create)
	router.DELETE("/vaga/:id", vagaController.Delete)
}
