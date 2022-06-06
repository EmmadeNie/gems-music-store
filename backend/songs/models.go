package songs

import (
	"github.com/emmadenie/gems-music-store/backend/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Song struct {
	gorm.Model
	Name   string `gorm: "" json: "name"`
	Artist string `json: "artist"`
	Album  string `json: "artist"`
	// ReleaseDate string `json: "releaseData`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Song{})
}

func (s *Song) CreateSong() *Song {
	db.NewRecord(s)
	db.Create(&s)

	return s
}
