package main

import (
	"log"
	"net/http"
	"os"

	_ "project-websocket/database"
	"project-websocket/routes"
)

func main() {
	router := routes.NewRouter()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:" + port, router))
}