package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateApartamento(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Apartamento criado com sucesso!"})
}
