package main

import (
	"blog-management/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)


func main() {
	cfg, _ := config.LoadConfig()
	
	log.Printf("Database configuration: %s:%s@%s/%s", 
        cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Name)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go Blog Management",
			"status": "running",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	address := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	if err := router.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}