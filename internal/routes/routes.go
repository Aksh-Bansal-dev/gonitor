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
	http.HandleFunc("/cpu", cpuHandler)
	http.HandleFunc("/ram", ramHandler)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Message: "pong"})
}

func ramHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
    res,err:= stat.GetRAMUsage()
    if err!=nil{
		log.Println(err)
		http.Error(w, "Internal server error", 500)
    }
	data := fmt.Sprintf("%.2f",res)
	json.NewEncoder(w).Encode(map[string]string{"data": data})
}
func cpuHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
    res,err := stat.GetCpuUsage()
	data := fmt.Sprintf("%.2f",res)
    if err!=nil{
		log.Println(err)
		http.Error(w, "Internal server error", 500)
        return
    }
	json.NewEncoder(w).Encode(map[string]string{"data": data})
}
