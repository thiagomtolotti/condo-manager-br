package vagaController

import (
	"backend/errs"
	vagaModel "backend/models/vaga"
	"backend/schemas"
	apartamentoService "backend/services/apartamento"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Create(c *gin.Context) {
	var body schemas.Vaga
	tryId := c.Param("apartamento_id")

	id, err := uuid.Parse(tryId)
	if err != nil {
		errs.BadRequestError(c, "id inválido")
		return
	}

	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		errs.BadRequestError(c, "Campos inválidos")
		return
	}

	exists, err := apartamentoService.Exists(id)
	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	if !exists {
		errs.BadRequestError(c, "Não existe apartamento com esse id")
		return
	}

	// TODO: Check if vaga with given number exists
	vagaId, err := vagaModel.Create(id, body)

	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "Vaga criada com sucesso",
			"id":      vagaId,
		},
	)
}
