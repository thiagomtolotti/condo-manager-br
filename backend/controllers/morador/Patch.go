package moradorController

import (
	moradorModel "backend/models/morador"
	"backend/schemas"
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

	// TODO: Check if new apartamento_id exists
	// TODO: Validate name size
	// TODO: Validate phone size

	err := moradorModel.Patch(cpf, body)

	if err != nil {
		fmt.Println("Error updating morador:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Morador editado com sucesso"})

}
