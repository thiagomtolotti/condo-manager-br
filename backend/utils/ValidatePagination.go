package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginatedParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func ValidatePagination(c *gin.Context) (PaginatedParams, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, sizeErr := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if err != nil || sizeErr != nil {
		return PaginatedParams{}, fmt.Errorf("invalid parameters")
	}

	if page < 1 || pageSize < 1 {
		return PaginatedParams{}, fmt.Errorf("invalid parameters")
	}

	return PaginatedParams{
		Page:     page,
		PageSize: pageSize,
	}, nil
}
