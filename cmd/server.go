package main

import "net/http"

type Server struct {
	Router	*http.ServeMux
}

func NewServer(router *http.ServeMux) *Server {
	return &Server{Router: router}
}
