package moradorController

import (
	"backend/errs"
	moradorModel "backend/models/morador"
	"backend/schemas"
	apartamentoService "backend/services/apartamento"
	"backend/utils/cpf"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Patch(c *gin.Context) {
	var body schemas.MoradorWithoutCPF
	tryCPF := c.Param("cpf")

	if err := c.ShouldBindJSON(&body); err != nil {
		errs.BadRequestError(c, "Campos inválidos")
		return
	}

	cpf, err := cpf.New(tryCPF)
	if err != nil {
		errs.BadRequestError(c, "CPF Inválido")
		return
	}

	exists, err := apartamentoService.Exists(body.Apartamento_id)

	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	if !exists {
		errs.BadRequestError(c, "Não existe um apartamento com esse id cadastrado")
		return
	}

	if len(body.Nome) > 100 {
		errs.BadRequestError(c, "O nome do morador deve ter no máximo 100 digitos")
		return
	}

	// TODO: Validate if phone only has numbers, spaces and dashes (Regex)
	if len(body.Telefone) > 15 {
		errs.BadRequestError(c, "O telefone deve ter no máximo 15 digitos")
		return
	}

	err = moradorModel.Patch(cpf, body)

	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Morador editado com sucesso",
		},
	)

}
