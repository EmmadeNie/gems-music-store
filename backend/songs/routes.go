package songs

import (
	"github.com/gorilla/mux"
)

func RegisterSongsRoutes(router *mux.Router) {
	router.HandleFunc("/song/", CreateSong).Methods("POST")
	router.HandleFunc("/song/", GetSong).Methods("GET")
	router.HandleFunc("/song/{songId}", GetSongById).Methods("GET")
	router.HandleFunc("/song/{songId}", UpdateSong).Methods("PUT")
	router.HandleFunc("/song/{songId}", DeleteSong).Methods("DELETE")
}
