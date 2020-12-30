package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	myServer := NewTodoServer(router)
	myServer.routes()

	log.Fatal(http.ListenAndServe(":8080", myServer.Router))
}
