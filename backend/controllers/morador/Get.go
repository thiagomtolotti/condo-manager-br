package moradorController

import (
	"backend/errs"
	moradorModel "backend/models/morador"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	params, err := utils.ValidatePagination(c)
	if err != nil {
		errs.HandleError(c, err)
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

	c.JSON(
		http.StatusOK,
		gin.H{
			"moradores":   data,
			"total_count": total,
		},
	)
}
