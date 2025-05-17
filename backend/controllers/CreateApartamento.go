package controllers

import (
	"backend/models"
	"backend/schemas"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateApartamento(c *gin.Context) {
	var body schemas.Apartamento

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Requisição Inválida"})
		return
	}

	err := models.CreateApartamento(body)

	if err != nil {
		fmt.Println("Erro creating apartment: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Apartamento criado com sucesso!"})
}
