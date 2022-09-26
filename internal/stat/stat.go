package stat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	res, err := GetRAMUsage()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
	}
	ram := fmt.Sprintf("%.2f%%", res)
	res, err = GetCpuUsage()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
	}
	cpu := fmt.Sprintf("%.2f%%", res)
	res, err = GetDiskUsage()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
	}
	disk := fmt.Sprintf("%.2f%%", res)
	json.NewEncoder(w).Encode(map[string]string{"ram": ram, "cpu": cpu, "disk": disk})
}
