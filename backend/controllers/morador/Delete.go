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
		errs.HandleError(c, errs.BadRequest("CPF inválido", err))
		return
	}

	appErr := moradorModel.Delete(cpf)

	if appErr != nil {
		errs.HandleError(c, appErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Morador excluído com sucesso"})
}
