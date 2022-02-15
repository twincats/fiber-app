package models

import (
	"github.com/twincats/fiber-app/server/database"
)

type Manga struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	StatusEnd string `json:"status_end"`
}

func (b *Manga) TableName() string {
	return "manga"
}

func GetMangas() []Manga {
	db := database.DBConn
	var manga []Manga
	db.Find(&manga)
	return manga
}
