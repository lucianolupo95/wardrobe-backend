package models

import "time"

type Clothing struct {
    ID           int       `db:"id" json:"id"`
    Name         string    `db:"name" json:"name"`
    PhotoURL     string    `db:"photo_url" json:"photoUrl"`
    SeasonID     int       `db:"season_id" json:"seasonId"`
    CategoryID   int       `db:"category_id" json:"categoryId"`
    StatusID     int       `db:"status_id" json:"statusId"`
    Visible      bool      `db:"visible" json:"visible"`
    CreationDate time.Time `db:"creation_date" json:"creationDate"`
    UpdatedAt    time.Time `db:"updated_at" json:"updatedAt"`
    Notes        string    `db:"notes" json:"notes"`
    DeletedAt   *time.Time `db:"deleted_at" json:"-"`
}
type ClothingCreateInput struct {
    Name       string `json:"name" binding:"required"`
    PhotoURL   string `json:"photoUrl" binding:"required"`
    SeasonID   int    `json:"seasonId" binding:"required"`
    CategoryID int    `json:"categoryId" binding:"required"`
    StatusID   int    `json:"statusId" binding:"required"`
    Visible    bool   `json:"visible"`
    Notes      string `json:"notes"`
}
