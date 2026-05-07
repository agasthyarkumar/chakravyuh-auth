package main

import (
	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.ConnectDB()

	router := gin.Default()

	routes.SetupRoutes(router)

	log.Println("Server running on :" + config.AppConfig.Port)

	router.Run(":" + config.AppConfig.Port)
}