package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucianolupo95/wardrobe-backend/handlers"
)

func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")

    clothes := api.Group("/clothes")
    {
        clothes.GET("/", handlers.GetAllClothes)
        clothes.GET("/:id", handlers.GetClothingByID)
        clothes.POST("/", handlers.CreateClothing)
        clothes.PUT("/:id", handlers.UpdateClothing)
        clothes.DELETE("/:id", handlers.DeleteClothing)
        clothes.PATCH("/:id/restore", handlers.RestoreClothing)
    }
}
