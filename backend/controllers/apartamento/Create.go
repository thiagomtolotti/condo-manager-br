package apartamentoController

import (
	"backend/errs"
	apartamentoModel "backend/models/apartamento"
	"backend/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var body schemas.Apartamento

	if err := c.ShouldBindJSON(&body); err != nil {
		var err = errs.BadRequest("Requisição inválida", err)
		errs.HandleError(c, err)
		return
	}

	id, err := apartamentoModel.CreateApartamento(body)
	if err != nil {
		errs.HandleError(c, err)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "Apartamento criado com sucesso!",
			"id":      id,
		},
	)
}
