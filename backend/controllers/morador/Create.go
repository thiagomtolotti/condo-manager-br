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

func Create(c *gin.Context) {
	var body schemas.Morador

	// TODO: Validate requests body on middleware
	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Campos Inválidos"})
		return
	}

	// TODO: Validate if no morador with the cpf exists
	if !utils.ValidateCPF(body.Cpf) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "CPF Inválido"})
		return
	}

	validApartment, queryErr := apartamentoService.Exists(body.Apartamento_id)

	if queryErr != nil {
		fmt.Println("Error querying apartment: ", queryErr)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	if !validApartment {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Apartamento inválido"})
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

	err := moradorModel.Create(body)

	if err != nil {
		fmt.Println("Error creating user: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	// TODO: Return resident id
	c.JSON(http.StatusCreated, gin.H{"message": "Morador criado com sucesso"})
}
