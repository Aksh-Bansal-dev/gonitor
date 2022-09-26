package config

import (
	"encoding/json"
	"log"
	"os"
)

type Endpoint struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Status string `json:"status"`
}
type Config struct {
	Endpoints []Endpoint `json:"endpoints"`
}

func ReadConfig() Config {
	data, err := os.ReadFile("gonitor.json")
	if err != nil {
		log.Println(err)
		return Config{}
	}
	var config Config
	err = json.Unmarshal(data, &config)
	return config
}
