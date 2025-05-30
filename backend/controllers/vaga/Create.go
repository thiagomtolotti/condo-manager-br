package vagaController

import (
	"backend/errs"
	apartmentoModel "backend/models/apartamento"
	vagaModel "backend/models/vaga"
	"backend/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Create(c *gin.Context) {
	var body schemas.Vaga
	id := c.Param("apartamento_id")
	apartamento_id, err := uuid.Parse(id)
	if err != nil {
		errs.HandleError(c, errs.BadRequest("id inválido", err))
		return
	}

	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		errs.HandleError(c, errs.BadRequest("campos inválidos", err))
		return
	}

	apartamento, appErr := apartmentoModel.FindById(apartamento_id)
	if appErr != nil {
		errs.HandleError(c, appErr)
		return
	}
	if apartamento == nil {
		errs.HandleError(c, errs.BadRequest("Não há apartamento com o id fornecido", nil))
		return
	}

	vaga, appErr := vagaModel.FindByNumber(body.Numero)
	if appErr != nil {
		errs.HandleError(c, appErr)
		return
	}
	if vaga != nil {
		errs.HandleError(c, errs.BadRequest("Já existe uma vaga com esse número", nil))
		return
	}

	vagaId, appErr := vagaModel.Create(apartamento_id, body)
	if appErr != nil {
		errs.HandleError(c, appErr)
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
