package apartamentoController

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	if len(strings.TrimSpace(id)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "É necessário fornecer um id válido"})
	}

	fmt.Println("Deleting apartment with id ", id)
}
