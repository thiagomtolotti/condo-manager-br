package apartamentoController

import (
	"backend/errs"
	apartamentoModel "backend/models/apartamento"
	moradorModel "backend/models/morador"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	if !utils.ValidateId(id) {
		errs.HandleError(c, errs.BadRequest("id inválido", nil))
		return
	}

	parsedId, _ := uuid.Parse(id)

	has_morador, err := moradorModel.HasMoradorInApartamento(parsedId)
	if err != nil {
		errs.HandleError(c, err)
		return
	}

	if has_morador {
		errs.HandleError(c, errs.BadRequest("Há moradores no apartamento", nil))
		return
	}

	// TODO: Check if apartamento has vagas (if it has throws an error on deleting)

	err = apartamentoModel.Delete(parsedId)
	if err != nil {
		errs.HandleError(c, err)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Apartamento excluído com sucesso!",
		},
	)
}
