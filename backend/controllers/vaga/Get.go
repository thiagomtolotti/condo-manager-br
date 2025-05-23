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

	vagas, err := vagaModel.Get(data.Page, data.PageSize)

	if err != nil {
		fmt.Println("Error fetching vagas:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	count, err := vagaModel.GetCount()
	if err != nil {
		fmt.Println("Error fetching vagas count:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"vagas":       vagas,
			"total_count": count,
		},
	)
}
