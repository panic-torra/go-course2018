package handlers

import (
	"encoding/json"
	"net/http"
)

func HandleListRouter(w http.ResponseWriter, _ *http.Request) {
	renderVideoArrayToJson(w, GetVideoArray())
}

func renderVideoArrayToJson(w http.ResponseWriter, videoArray []VideoInfo) {
	b, err := json.Marshal(videoArray)
	handleJsonRender(w, b, err)
}
