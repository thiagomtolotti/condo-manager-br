package moradorController

import (
	moradorModel "backend/models/morador"
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

	data, err := moradorModel.Get(params.Page, params.PageSize)

	if err != nil {
		fmt.Println("Error fetching morador: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	total, err := moradorModel.GetCount()
	if err != nil {
		fmt.Println("Error fetching morador total count:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	// TODO: Return the total count of moradores (for pagination in the FE)
	c.JSON(
		http.StatusOK,
		gin.H{
			"moradores":   data,
			"total_count": total,
		},
	)
}
