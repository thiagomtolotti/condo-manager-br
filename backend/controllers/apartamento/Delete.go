package apartamentoController

import (
	"backend/errs"
	apartamentoModel "backend/models/apartamento"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Delete(c *gin.Context) {
	tryId := c.Param("id")

	id, err := uuid.Parse(tryId)

	if err != nil {
		errs.BadRequestError(c, "id inválido")
		return
	}

	// TODO: Check if apartamento has moradores (if it has throws an error on deleting)
	// TODO: Check if apartamento has vagas (if it has throws an error on deleting)

	err = apartamentoModel.Delete(id)

	if err != nil {
		if errors.Is(err, apartamentoModel.ErrNotFound) {
			errs.BadRequestError(c, "id inválido")
			return
		}

		errs.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Apartamento excluído com sucesso!"})
}
