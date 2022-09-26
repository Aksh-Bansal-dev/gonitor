package healthcheck

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"example.com/gonitor/internal/config"
)

func pingServer(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	if err != nil || res.StatusCode != 200 {
		return errors.New("Service down!")
	}
	return nil
}

func HealthHandler(w http.ResponseWriter, r *http.Request, endpoints []config.Endpoint) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if endpoints == nil {
		log.Println("Error parsing config")
		json.NewEncoder(w).Encode("No endpoints.")
		return
	}

	for i, endpoint := range endpoints {
		err := pingServer(endpoint.URL)
		if err != nil {
			endpoints[i].Status = "down"
		} else {
			endpoints[i].Status = "up"
		}
	}

	json.NewEncoder(w).Encode(endpoints)
}
