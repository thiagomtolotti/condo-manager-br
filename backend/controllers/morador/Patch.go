package moradorController

import (
	"backend/errs"
	moradorModel "backend/models/morador"
	"backend/schemas"
	apartamentoService "backend/services/apartamento"
	"backend/utils/cpf"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Patch(c *gin.Context) {
	var body schemas.MoradorWithoutCPF
	try := c.Param("cpf")
	cpf, err := cpf.New(try)

	if err != nil {
		errs.HandleError(c, errs.BadRequest("CPF inválido", err))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		errs.HandleError(c, errs.BadRequest("campos inválidos", err))
		return
	}

	appErr := apartamentoService.Exists(body.Apartamento_id)
	if appErr != nil {
		errs.HandleError(c, appErr)
		return
	}

	if len(body.Nome) > 100 {
		errs.HandleError(c, errs.BadRequest("O nome do morador deve ter no máximo 100 digitos", nil))
		return
	}

	// TODO: Validate if phone only has numbers, spaces and dashes (Regex)
	if len(body.Telefone) > 15 {
		errs.HandleError(c, errs.BadRequest("O telefone deve ter no máximo 15 digitos", nil))
		return
	}

	appErr = moradorModel.Patch(cpf, body)
	if err != nil {
		fmt.Println("Error updating morador:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Morador editado com sucesso"})

}
