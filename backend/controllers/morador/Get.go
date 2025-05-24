package moradorController

import (
	"backend/errs"
	moradorModel "backend/models/morador"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	params, err := utils.ValidatePagination(c)

	if err != nil {
		errs.BadRequestError(c, "Parâmetros inválidos")
		return
	}

	data, err := moradorModel.Get(params.Page, params.PageSize)

	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	total, err := moradorModel.GetCount()
	if err != nil {
		errs.InternalServerError(c, err)
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
