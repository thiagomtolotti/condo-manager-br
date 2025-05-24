package vagaController

import (
	"backend/errs"
	vagaModel "backend/models/vaga"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	data, err := utils.ValidatePagination(c)

	if err != nil {
		errs.BadRequestError(c, "Parâmetros inválidos")
		return
	}

	vagas, err := vagaModel.Get(data.Page, data.PageSize)
	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	count, err := vagaModel.GetCount()
	if err != nil {
		errs.InternalServerError(c, err)
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
