package moradorController

import (
	moradorModel "backend/models/morador"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	cpf := c.Param("cpf")

	if !utils.ValidateCPF(cpf) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "CPF Inválido"})
		return
	}

	success, err := moradorModel.Delete(cpf)

	if err != nil {
		fmt.Println("Error deleting morador: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"message": "CPF Inválido"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Morador excluído com sucesso"})
}
