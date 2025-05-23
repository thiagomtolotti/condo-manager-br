package moradorController

import (
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
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Campos Inválidos"})
		return
	}

	cpf, err := cpf.New(body.Cpf)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "CPF Inválido"})
	}

	isValid, err := moradorService.Validate(cpf)

	if err != nil {
		fmt.Println("Error validating CPF:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "CPF Inválido ou com morador já cadastrado"})
		return
	}

	validApartment, err := apartamentoService.Exists(body.Apartamento_id)

	if err != nil {
		fmt.Println("Error querying apartment: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
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

	err = moradorModel.Create(body)

	if err != nil {
		fmt.Println("Error creating user: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Morador criado com sucesso"})
}
