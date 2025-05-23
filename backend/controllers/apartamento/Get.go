package apartamentoController

import (
	apartamentoModel "backend/models/apartamento"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	params, err := utils.ValidatePagination(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	rows, err := apartamentoModel.GetApartamento(params.Page, params.PageSize)

	if err != nil {
		fmt.Println("Error fetching apartment: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	total, err := apartamentoModel.GetCount()

	if err != nil {
		fmt.Println("Error fetching apartment count:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"apartamentos": rows,
			"total_count":  total,
		},
	)
}
