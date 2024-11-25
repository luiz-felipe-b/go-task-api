package main

import (
	"to-do-api/config"
	"to-do-api/internal/tasks"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load the configuration
	cfg := config.LoadConfig()

	// Create a new gin router
	router := gin.Default()

	// Create a new task service
	tasks.RegisterRoutes(router)

	// Start the server
	router.Run(":" + cfg.Port)
}
