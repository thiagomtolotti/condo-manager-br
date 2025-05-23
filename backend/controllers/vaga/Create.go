package vagaController

import (
	vagaModel "backend/models/vaga"
	"backend/schemas"
	apartamentoService "backend/services/apartamento"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Create(c *gin.Context) {
	var body schemas.Vaga
	id := c.Param("apartamento_id")
	apartamento_id, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Apartamento inválido"})
		return
	}

	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Requisição inválida"})
		return
	}

	exists, apid_error := apartamentoService.Exists(apartamento_id)

	if apid_error != nil {
		fmt.Println("Error checking if apartment exists: ", apid_error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Apartamento inválido"})
		return
	}

	// TODO: Check if vaga with given number exists
	vagaId, err := vagaModel.Create(apartamento_id, body)

	if err != nil {
		fmt.Println("Error creating parking space:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "Vaga criada com sucesso",
			"id":      vagaId,
		},
	)
}
