package apartamentoController

import (
	apartmentModel "backend/models/apartamento"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	if !utils.ValidateId(id) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "É necessário fornecer um id válido"})
		return
	}

	parsedId, _ := uuid.Parse(id)

	success, err := apartmentModel.Delete(parsedId)

	if err != nil {
		fmt.Println("Error deleting apartment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Id inválido"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Apartamento excluído com sucesso!"})
}
