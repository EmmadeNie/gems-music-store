package songs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/emmadenie/gems-music-store/backend/utils"
	"github.com/gorilla/mux"
)

func GetSong(w http.ResponseWriter, r *http.Request) {
	newSongs := GetAllSongs()

	res, _ := json.Marshal(newSongs)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetSongById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songId := vars["songId"]

	ID, err := strconv.ParseInt(songId, 0, 0)

	if err != nil {
		fmt.Println("err while parsing")
	}

	songDetails, _ := GetSongByIdModel(ID)
	res, _ := json.Marshal(songDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateSong(w http.ResponseWriter, r *http.Request) {
	CreateSong := &Song{}
	utils.ParseBody(r, CreateSong)
	s := CreateSong.CreateSong()
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songId := vars["songId"]
	ID, err := strconv.ParseInt(songId, 0, 0)
	if err != nil {
		fmt.Println("err while parsing")
	}
	song := DeleteSongModel(ID)

	res, _ := json.Marshal(song)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateSong(w http.ResponseWriter, r *http.Request) {
	var updateSong = &Song{}
	utils.ParseBody(r, updateSong)
	vars := mux.Vars(r)
	songId := vars["songId"]

	ID, err := strconv.ParseInt(songId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}
	songDetails, db := GetSongByIdModel(ID)
	if updateSong.Name != "" {
		songDetails.Name = updateSong.Name
	}
	if updateSong.Name != "" {
		songDetails.Name = updateSong.Name
	}
	if updateSong.Artist != "" {
		songDetails.Artist = updateSong.Artist
	}
	if updateSong.Album != "" {
		songDetails.Album = updateSong.Album
	}
	db.Save(&songDetails)
	res, _ := json.Marshal(songDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
