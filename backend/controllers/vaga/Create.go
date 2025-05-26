package vagaController

import (
	"backend/errs"
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
		errs.HandleError(c, errs.BadRequest("id inválido", err))
		return
	}

	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		errs.HandleError(c, errs.BadRequest("campos inválidos", err))
		return
	}

	appErr := apartamentoService.Exists(apartamento_id)
	if appErr != nil {
		errs.HandleError(c, appErr)
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
