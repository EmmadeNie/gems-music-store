package songs

import (
	"github.com/gorilla/mux"
)

func RegisterSongsRoutes(router *mux.Router) {
	router.HandleFunc("/song/", CreateSong).Methods("POST")

}
