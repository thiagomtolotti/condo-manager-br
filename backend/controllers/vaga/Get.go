package vagaController

import (
	vagaModel "backend/models/vaga"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	data, err := utils.ValidatePagination(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Parâmetros inválidos"})
	}

	vagas, queryErr := vagaModel.Get(data.Page, data.PageSize)

	if queryErr != nil {
		fmt.Println("Error fetching vagas:", queryErr)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	// TODO: Return the total count of parking spaces (for pagination in the FE)
	c.JSON(http.StatusOK, gin.H{"vagas": vagas})
}
