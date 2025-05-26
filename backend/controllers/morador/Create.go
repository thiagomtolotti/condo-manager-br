package moradorController

import (
	"backend/errs"
	moradorModel "backend/models/morador"
	"backend/schemas"
	apartamentoService "backend/services/apartamento"
	moradorService "backend/services/morador"
	"backend/utils/cpf"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var body schemas.Morador

	// TODO: Validate requests body on middleware
	if err := c.ShouldBindJSON(&body); err != nil {
		errs.HandleError(c, errs.BadRequest("campos inválidos", err))
		return
	}

	cpf, err := cpf.New(body.Cpf)
	if err != nil {
		errs.HandleError(c, errs.BadRequest("CPF inválido", err))
		return
	}

	appErr := moradorService.Validate(cpf)
	if appErr != nil {
		errs.HandleError(c, appErr)
		return
	}

	appErr = apartamentoService.Exists(body.Apartamento_id)
	if appErr != nil {
		errs.HandleError(c, appErr)
		return
	}

	if len(body.Nome) > 100 {
		errs.HandleError(c, errs.BadRequest("O nome do morador deve ter no máximo 100 digitos", err))
		return
	}

	// TODO: Validate if phone only has numbers, spaces and dashes (Regex)
	if len(body.Telefone) > 15 {
		errs.HandleError(c, errs.BadRequest("O telefone deve ter no máximo 15 dígitos", nil))
		return
	}

	appErr = moradorModel.Create(body)
	if err != nil {
		fmt.Println("Error creating user: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Morador criado com sucesso"})
}
