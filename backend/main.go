package main

import (
	"fmt"
	"linkedin/database"
	"linkedin/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//	@title			LinkedIn API
//	@version		1.0
//	@description	Backend API endpoints for LinkedIn.

//	@host		software.ardapektezol.com/api
//	@BasePath	/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	database.ConnectDB()
	routes.SetupRoutes(router)
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
