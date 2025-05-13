package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	router := gin.Default()

	fmt.Println("Starting server on port " + PORT)
	router.Run(":" + PORT)
}
