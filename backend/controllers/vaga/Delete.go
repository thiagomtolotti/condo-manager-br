package vagaController

import (
	"backend/errs"
	vagaModel "backend/models/vaga"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)

	if err != nil {
		errs.BadRequestError(c, "Id inválido")
		return
	}

	success, err := vagaModel.Delete(uuid)

	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	if !success {
		errs.BadRequestError(c, "id inválido")
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Vaga excluída com sucesso",
		},
	)

}
