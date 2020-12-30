package main

func (s *Server) routes() {
	s.Router.HandleFunc("/todos/:id", TodoShow)
	s.Router.HandleFunc("/todos", TodoIndex)
	s.Router.HandleFunc("/", Index)
}
