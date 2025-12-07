package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllClothes(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "get all clothes"})
}

func GetClothingByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"message": "get clothing by id", "id": id})
}

func CreateClothing(c *gin.Context) {
    c.JSON(http.StatusCreated, gin.H{"message": "create clothing"})
}

func UpdateClothing(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"message": "update clothing", "id": id})
}

func DeleteClothing(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"message": "delete clothing", "id": id})
}
