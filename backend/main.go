package main

import (
	routes "backend/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	router := gin.Default()

	routes.RegisterRoutes(router)

	fmt.Println("Starting server on port " + PORT)
	router.Run(":" + PORT)
}
