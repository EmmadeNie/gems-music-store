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

func GetAllSongs() []Song {
	var Songs []Song
	db.Find(&Songs)
	return Songs
}

func GetSongByIdModel(Id int64) (*Song, *gorm.DB) {
	var getSong Song
	db := db.Where("ID=?", Id).Find(&getSong)
	return &getSong, db
}

func DeleteSongModel(ID int64) Song {
	var song Song
	db.Where("ID=?", ID).Delete(song)
	return song

}
