package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	myServer := NewServer(router)
	myServer.routes()

	log.Fatal(http.ListenAndServe(":8080", router))
}
