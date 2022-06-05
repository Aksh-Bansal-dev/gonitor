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

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	log.Println("Server started at port:", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
