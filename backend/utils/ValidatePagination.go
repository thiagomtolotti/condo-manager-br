package utils

import (
	"backend/errs"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginatedParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func ValidatePagination(c *gin.Context) (PaginatedParams, *errs.AppError) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, sizeErr := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if err != nil || sizeErr != nil {
		return PaginatedParams{}, errs.BadRequest("parâmetros inválidos", nil)
	}

	if page < 1 || pageSize < 1 {
		return PaginatedParams{}, errs.BadRequest("parâmetros inválidos", nil)
	}

	return PaginatedParams{
		Page:     page,
		PageSize: pageSize,
	}, nil
}
