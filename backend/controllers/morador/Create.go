package moradorController

import (
	"backend/errs"
	moradorModel "backend/models/morador"
	"backend/schemas"
	apartamentoService "backend/services/apartamento"
	moradorService "backend/services/morador"
	"backend/utils/cpf"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var body schemas.Morador

	// TODO: Validate requests body on middleware
	if err := c.ShouldBindJSON(&body); err != nil {
		errs.BadRequestError(c, "Campos inválidos")
		return
	}

	cpf, err := cpf.New(body.Cpf)
	if err != nil {
		errs.BadRequestError(c, "CPF inválido")
		return
	}

	isValid, err := moradorService.Validate(cpf)
	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "CPF Inválido ou com morador já cadastrado"})
		return
	}

	validApartment, err := apartamentoService.Exists(body.Apartamento_id)
	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	if !validApartment {
		errs.BadRequestError(c, "Apartamento inválido")
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

	err = moradorModel.Create(body)

	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "Morador criado com sucesso",
		},
	)
}
