package apartamentoController

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	if !validateId(id) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "É necessário fornecer um id válido"})
		return
	}

	fmt.Println("Deleting apartment with id ", id)
}

func validateId(id string) bool {
	if len(strings.TrimSpace(id)) == 0 {
		return false
	}

	if err := uuid.Validate(id); err != nil {
		return false
	}

	return true
}
