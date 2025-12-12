package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucianolupo95/wardrobe-backend/db"
	"github.com/lucianolupo95/wardrobe-backend/models"
)

func GetAllClothes(c *gin.Context) {
    clothes := []models.Clothing{}

    err := db.DB.Select(&clothes, "SELECT * FROM clothes WHERE deleted_at IS NULL ORDER BY id")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, clothes)
}



func GetClothingByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid id",
        })
        return
    }

    var item models.Clothing
    query := "SELECT * FROM clothes WHERE id = $1 AND deleted_at IS NULL LIMIT 1"

    err = db.DB.Get(&item, query, id)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "clothing item not found",
            })
            return
        }

        c.JSON(http.StatusInternalServerError, gin.H{
            "error":  "database error",
            "detail": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, item)
}


func CreateClothing(c *gin.Context) {
    var input models.ClothingCreateInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid json",
            "detail": err.Error(),
        })
        return
    }

    query := `
        INSERT INTO clothes
            (name, photo_url, season_id, category_id, status_id, visible, notes)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, creation_date, updated_at
    `

    var created models.Clothing
    err := db.DB.QueryRow(
        query,
        input.Name,
        input.PhotoURL,
        input.SeasonID,
        input.CategoryID,
        input.StatusID,
        input.Visible,
        input.Notes,
    ).Scan(
        &created.ID,
        &created.CreationDate,
        &created.UpdatedAt,
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":  "database insert error",
            "detail": err.Error(),
        })
        return
    }

    created.Name = input.Name
    created.PhotoURL = input.PhotoURL
    created.SeasonID = input.SeasonID
    created.CategoryID = input.CategoryID
    created.StatusID = input.StatusID
    created.Visible = input.Visible
    created.Notes = input.Notes

    c.JSON(http.StatusCreated, created)
}

func UpdateClothing(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    var input models.ClothingCreateInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":  "invalid json",
            "detail": err.Error(),
        })
        return
    }

    query := `
        UPDATE clothes
        SET
            name = $1,
            photo_url = $2,
            season_id = $3,
            category_id = $4,
            status_id = $5,
            visible = $6,
            notes = $7,
            updated_at = NOW()
        WHERE id = $8
        RETURNING id, creation_date, updated_at
    `

    var updated models.Clothing
    err = db.DB.QueryRow(
        query,
        input.Name,
        input.PhotoURL,
        input.SeasonID,
        input.CategoryID,
        input.StatusID,
        input.Visible,
        input.Notes,
        id,
    ).Scan(
        &updated.ID,
        &updated.CreationDate,
        &updated.UpdatedAt,
    )

    if err == sql.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
        return
    }
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    updated.Name = input.Name
    updated.PhotoURL = input.PhotoURL
    updated.SeasonID = input.SeasonID
    updated.CategoryID = input.CategoryID
    updated.StatusID = input.StatusID
    updated.Visible = input.Visible
    updated.Notes = input.Notes

    c.JSON(http.StatusOK, updated)
}


func DeleteClothing(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    query := `
        UPDATE clothes
        SET deleted_at = NOW()
        WHERE id = $1 AND deleted_at IS NULL
    `

    res, err := db.DB.Exec(query, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    affected, _ := res.RowsAffected()
    if affected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "item not found or already deleted"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"deleted": id})
}
func RestoreClothing(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    query := `
        UPDATE clothes
        SET deleted_at = NULL, updated_at = NOW()
        WHERE id = $1 AND deleted_at IS NOT NULL
    `

    res, err := db.DB.Exec(query, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    affected, _ := res.RowsAffected()
    if affected == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "item not found or not deleted",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "restored": id,
    })
}


