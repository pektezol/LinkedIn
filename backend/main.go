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

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	// router.Use(cors.Default())
	database.ConnectDB()
	routes.SetupRoutes(router)
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
