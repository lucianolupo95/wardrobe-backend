package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucianolupo95/wardrobe-backend/db"
	"github.com/lucianolupo95/wardrobe-backend/routes"
)

func main() {
    db.Connect()

    r := gin.Default()

    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // Registrar rutas
    routes.RegisterRoutes(r)

    r.Run(":8080")
}
