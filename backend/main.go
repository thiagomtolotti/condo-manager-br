package main

import (
	"backend/db"
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

	conn := utils.ConnectToDatabase()
	db.Migrate(conn)
	fmt.Println("Starting server on port " + PORT)

	fmt.Print("\n")

	router.Run(":" + PORT)
}
