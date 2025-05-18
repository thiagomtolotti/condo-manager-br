package moradorController

import (
	moradorModel "backend/models/morador"
	"backend/schemas"
	apartamentoService "backend/services/apartamento"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Patch(c *gin.Context) {
	var body schemas.MoradorWithoutCPF
	cpf := c.Param("cpf")

	if !utils.ValidateCPF(cpf) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "CPF Inválido"})
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Campos Inválidos"})
		return
	}

	exists, existsErr := apartamentoService.Exists(body.Apartamento_id)

	if existsErr != nil {
		fmt.Println("Error checking if apartamento exists:", existsErr)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Apartamento não existe"})
		return
	}

	if len(body.Nome) > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "O nome do morador deve ter no máximo 100 digitos"})
		return
	}

	// TODO: Validate if phone only has numbers, spaces and dashes (Regex)
	if len(body.Telefone) > 15 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "O telefone deve ter no máximo 15 digitos"})
		return
	}

	err := moradorModel.Patch(cpf, body)

	if err != nil {
		fmt.Println("Error updating morador:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Morador editado com sucesso"})

}
