package moradorController

import (
	"backend/errs"
	apartmentoModel "backend/models/apartamento"
	moradorModel "backend/models/morador"
	"backend/schemas"
	"backend/utils/cpf"
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

	morador, appErr := moradorModel.FindByCPF(cpf)
	if appErr != nil {
		errs.HandleError(c, appErr)
		return
	}
	if morador != nil {
		errs.HandleError(c, errs.BadRequest("Já existe um morador com este id", nil))
		return
	}

	apartamento, appErr := apartmentoModel.FindById(body.Apartamento_id)
	if appErr != nil {
		errs.HandleError(c, appErr)
		return
	}
	if apartamento == nil {
		errs.HandleError(c, errs.BadRequest("Não há apartamento com este id", nil))
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
	if appErr != nil {
		errs.HandleError(c, appErr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Morador criado com sucesso"})
}
