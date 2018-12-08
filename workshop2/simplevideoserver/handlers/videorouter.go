package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func HandleVideoRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	video, _ := GetVideoById(id)
	renderVideoToJson(w, video)
}

func renderVideoToJson(w http.ResponseWriter, video VideoInfo) {
	b, err := json.Marshal(video)
	handleJsonRender(w, b, err)
}
