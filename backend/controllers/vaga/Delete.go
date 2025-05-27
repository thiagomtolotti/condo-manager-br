package vagaController

import (
	"backend/errs"
	vagaModel "backend/models/vaga"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Delete(c *gin.Context) {
	try := c.Param("id")
	id, err := uuid.Parse(try)
	if err != nil {
		errs.HandleError(c, errs.BadRequest("id inválido", err))
		return
	}

	appErr := vagaModel.Delete(id)
	if appErr != nil {
		errs.HandleError(c, appErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vaga excluída com sucesso"})

}
