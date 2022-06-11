package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/gonitor/internal/stat"
)

type PingResponse struct {
	Message string `json:"message"`
}

var FileServer = http.FileServer(http.Dir("./static"))

func Routes() {
	http.Handle("/static/", http.StripPrefix("/static/", FileServer))
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/stats", statsHandler)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Message: "pong"})
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	res, err := stat.GetRAMUsage()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
	}
	ram := fmt.Sprintf("%.2f", res)
	res, err = stat.GetCpuUsage()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
	}
	cpu := fmt.Sprintf("%.2f", res)
	res, err = stat.GetDiskUsage()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
	}
	disk := fmt.Sprintf("%.2f", res)
	json.NewEncoder(w).Encode(map[string]string{"ram": ram, "cpu": cpu, "disk": disk})
}
