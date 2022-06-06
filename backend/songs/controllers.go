package songs

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emmadenie/gems-music-store/backend/utils"
)

func CreateSong(w http.ResponseWriter, r *http.Request) {
	fmt.Println("lala")
	CreateSong := &Song{}
	utils.ParseBody(r, CreateSong)
	s := CreateSong.CreateSong()
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
