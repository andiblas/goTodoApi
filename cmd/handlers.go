package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Mr. %q\n", html.EscapeString(request.URL.Path))
}

func TodoIndex(writer http.ResponseWriter, request *http.Request) {
	todos := Todos{
		Todo{
			Name: "Learn Go basics",
		},
		Todo{
			Name: "Learn routing",
		},
	}

	json.NewEncoder(writer).Encode(todos)
}

func TodoShow(writer http.ResponseWriter, request *http.Request) {
	todoId, err := strconv.Atoi(strings.TrimPrefix(request.URL.Path, "/todos/"))
	if err != nil {
		fmt.Fprint(writer, "Invalid TODO id.\n")
		return
	}

	fmt.Fprintf(writer, "Showing TODO id: %d\n", todoId)
}
