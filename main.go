package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lucianolupo95/wardrobe-backend/db"
	"github.com/lucianolupo95/wardrobe-backend/routes"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders: []string{"*"},
	}))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.Use(func(c *gin.Context) {
		fmt.Println("=> Nueva request")
		c.Next()
	})

	// Registrar rutas
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
