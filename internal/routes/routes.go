package routes

import (
	"encoding/json"
	"net/http"

	"example.com/gonitor/internal/config"
	"example.com/gonitor/internal/healthcheck"
	"example.com/gonitor/internal/stat"
)

type PingResponse struct {
	Message string `json:"message"`
}

var FileServer = http.FileServer(http.Dir("./static"))

func Routes() {
	configData := config.ReadConfig()
	http.Handle("/", http.StripPrefix("/", FileServer))
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/stats", stat.StatsHandler)
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		healthcheck.HealthHandler(w, r, configData.Endpoints)
	})
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Message: "pong"})
}
