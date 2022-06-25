package healthcheck

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type healthcheckRes struct {
	Message string `json:"message"`
}

type Endpoint struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Config struct {
	Endpoints Endpoint `json:"endpoints"`
}

func pingServer(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var v healthcheckRes
	err = json.Unmarshal([]byte(string(body)), &v)
	if err != nil || v.Message != "pong" {
		return errors.New("Service down!")
	}
	return nil
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	list := []Endpoint{{
		Name: "serviceA",
		URL:  "http://localhost:3000",
	}}

	json.NewEncoder(w).Encode(map[string]string{})
}
