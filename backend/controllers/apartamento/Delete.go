package apartamentoController

import (
	apartmentModel "backend/models/apartamento"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	if !validateId(id) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "É necessário fornecer um id válido"})
		return
	}

	parsedId, _ := uuid.Parse(id)

	err := apartmentModel.Delete(parsedId)

	if err != nil {
		fmt.Println("Error deleting apartment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Apartamento excluído com sucesso!"})
}

func validateId(id string) bool {
	if len(strings.TrimSpace(id)) == 0 {
		return false
	}

	if err := uuid.Validate(id); err != nil {
		return false
	}

	return true
}
