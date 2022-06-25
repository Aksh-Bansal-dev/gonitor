package routes

import (
	"encoding/json"
	"net/http"

	"example.com/gonitor/internal/healthcheck"
	"example.com/gonitor/internal/stat"
)

type PingResponse struct {
	Message string `json:"message"`
}

var FileServer = http.FileServer(http.Dir("./static"))

func Routes() {
	http.Handle("/static/", http.StripPrefix("/static/", FileServer))
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/stats", stat.StatsHandler)
	http.HandleFunc("/healthcheck", healthcheck.HealthHandler)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Message: "pong"})
}
