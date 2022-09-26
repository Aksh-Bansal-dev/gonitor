package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/gonitor/internal/routes"
)

func main() {
	routes.Routes()

	addr := os.Getenv("ADDR")
	port := os.Getenv("PORT")
	if len(addr) == 0 {
		addr = "localhost"
	}
	if len(port) == 0 {
		port = "5000"
	}
	log.Printf("Server started at %s:%s", addr, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), nil)
}
