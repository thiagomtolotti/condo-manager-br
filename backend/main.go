package main

import (
	routes "backend/router"
	"backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	router := gin.Default()

	utils.LoadEnvironment()
	routes.RegisterRoutes(router)

	fmt.Println()
	fmt.Println("Starting server on port " + PORT)
	fmt.Println()
	fmt.Println()
	router.Run(":" + PORT)
}
