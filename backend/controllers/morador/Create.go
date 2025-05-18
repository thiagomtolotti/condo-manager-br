package moradorController

import (
	moradorModel "backend/models/morador"
	"backend/schemas"
	moradorService "backend/services/morador"
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

	validCpf := moradorService.ValidateCPF(body.Cpf)

	if !validCpf {
		c.JSON(http.StatusBadRequest, gin.H{"message": "CPF Inválido"})
		return
	}

	// TODO: Validate apartamento_id
	// TODO: Validate nome length
	// TODO: Validate telefone

	err := moradorModel.Create(body)

	if err != nil {
		fmt.Println("Error creating user: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	// TODO: Return resident id
	c.JSON(http.StatusCreated, gin.H{"message": "Morador criado com sucesso"})
}
