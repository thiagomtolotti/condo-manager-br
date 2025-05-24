package moradorController

import (
	"backend/errs"
	moradorModel "backend/models/morador"
	"backend/utils/cpf"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	try := c.Param("cpf")
	cpf, err := cpf.New(try)

	if err != nil {
		errs.BadRequestError(c, "CPF inválido")
		return
	}

	success, err := moradorModel.Delete(cpf)
	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	if !success {
		errs.BadRequestError(c, "CPF inválido")
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Morador excluído com sucesso",
		},
	)
}
