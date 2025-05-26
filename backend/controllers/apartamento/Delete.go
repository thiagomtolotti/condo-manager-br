package apartamentoController

import (
	"backend/errs"
	apartamentoModel "backend/models/apartamento"
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

	err := apartamentoModel.Delete(parsedId)

	// TODO: Check if apartamento has moradores (if it has throws an error on deleting)
	// TODO: Check if apartamento has vagas (if it has throws an error on deleting)
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
