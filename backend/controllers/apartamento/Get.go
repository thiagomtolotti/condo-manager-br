package apartamentoController

import (
	apartmentModel "backend/models/apartamento"
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

	rows, fetchErr := apartmentModel.GetApartamento(params.Page, params.PageSize)

	if fetchErr != nil {
		fmt.Println("Error fetching apartment: ", fetchErr)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	// TODO: Return the total count of apartments (for pagination in the FE)
	c.JSON(http.StatusOK, gin.H{"apartamentos": rows})
}
