package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
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
	homePath, _ := os.UserHomeDir()
	data, err := os.ReadFile(path.Join(homePath, "gonitor.json"))
	if err != nil {
		log.Println(err)
		return Config{}
	}
	var config Config
	err = json.Unmarshal(data, &config)
	return config
}
