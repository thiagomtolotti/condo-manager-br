package vagaController

import (
	vagaModel "backend/models/vaga"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Id inválido"})
		return
	}

	success, queryErr := vagaModel.Delete(uuid)

	if queryErr != nil {
		fmt.Println("Error deleting parking space:", queryErr)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Id inválido"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vaga excluída com sucesso"})

}
