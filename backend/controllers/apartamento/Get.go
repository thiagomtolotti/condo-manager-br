package apartamentoController

import (
	"backend/errs"
	apartamentoModel "backend/models/apartamento"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	params, err := utils.ValidatePagination(c)

	if err != nil {
		errs.HandleError(c, errs.BadRequest("parâmetros inválidos", err))
		return
	}

	rows, err := apartamentoModel.GetApartamento(params.Page, params.PageSize)
	if err != nil {
		errs.HandleError(c, err)
		return
	}

	total, err := apartamentoModel.GetCount()
	if err != nil {
		errs.HandleError(c, err)
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
