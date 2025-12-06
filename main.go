package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucianolupo95/wardrobe-backend/db"
)

func main() {
    db.Connect()
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "ok",
        })
    })

    r.Run(":8080")
}
