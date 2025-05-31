package moradorController

import (
	"backend/errs"
	apartmentoModel "backend/models/apartamento"
	moradorModel "backend/models/morador"
	"backend/schemas"
	"backend/utils"
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

	apartamento, appErr := apartmentoModel.FindById(body.Apartamento_id)
	if appErr != nil {
		errs.HandleError(c, appErr)
		return
	}
	if apartamento == nil {
		errs.HandleError(c, errs.BadRequest("Não há apartamento com o id fornecido", nil))
		return
	}

	if len(body.Nome) > 100 {
		errs.HandleError(c, errs.BadRequest("O nome do morador deve ter no máximo 100 digitos", nil))
		return
	}

	if !utils.ValidatePhone(body.Telefone) {
		errs.HandleError(c, errs.BadRequest("Telefone inválido", nil))
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
