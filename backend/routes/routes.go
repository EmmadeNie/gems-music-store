package routes

import (
	"github.com/emmadenie/gems-music-store/backend/songs"
	"github.com/gorilla/mux"
)

var RegisterMusicStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/song/", songs.CreateSong).Methods("POST")

}
