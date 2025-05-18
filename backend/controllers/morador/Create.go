package moradorController

import (
	"backend/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var body schemas.Morador

	// TODO: Validate requests body on middleware
	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Campos Inválidos"})
	}

	fmt.Println("Creating user", body)
}
