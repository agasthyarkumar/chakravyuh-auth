package main

import (
	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/redis"
	"auth-service/internal/routes"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.ConnectDB()

	// Initialize Redis (non-blocking, gracefully handles connection failures)
	_ = redis.ConnectRedis(config.AppConfig.RedisURL)

	// Ensure Redis closes gracefully
	defer redis.Close()

	router := gin.Default()

	routes.SetupRoutes(router)

	// Graceful shutdown handling
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Println("Shutting down server...")

		_ = redis.Close()
		os.Exit(0)
	}()

	log.Println("Server running on :" + config.AppConfig.Port)

	router.Run(":" + config.AppConfig.Port)
}