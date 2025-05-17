package main

import (
	routes "backend/router"
	"backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	utils.LoadEnvironment()
	router := gin.Default()

	routes.RegisterRoutes(router)

	fmt.Print("\n")

	utils.ConnectToDatabase()
	fmt.Println("Starting server on port " + PORT)

	fmt.Print("\n")

	router.Run(":" + PORT)
}
