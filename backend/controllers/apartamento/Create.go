package apartamentoController

import (
	apartmentModel "backend/models/apartamento"
	"backend/schemas"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var body schemas.Apartamento

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Requisição Inválida"})
		return
	}

	err := apartmentModel.CreateApartamento(body)

	if err != nil {
		fmt.Println("Erro creating apartment: ", err)

		// TODO: Treat the error accordingly (HTTPError strategy)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	// TODO: Return id with the apartment created
	c.JSON(http.StatusOK, gin.H{"message": "Apartamento criado com sucesso!"})
}
