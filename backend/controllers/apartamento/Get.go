package apartamentoController

import (
	apartmentModel "backend/models/apartamento"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	params, err := validateParameters(c)

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

type PaginatedParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func validateParameters(c *gin.Context) (PaginatedParams, error) {
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
