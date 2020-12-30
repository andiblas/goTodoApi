package main

import "net/http"

type TodoServer struct {
	Router *http.ServeMux
}

func NewTodoServer(router *http.ServeMux) *TodoServer {
	return &TodoServer{Router: router}
}
