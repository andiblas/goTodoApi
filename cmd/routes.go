package main

func (s *TodoServer) routes() {
	s.Router.HandleFunc("/todos/:id", TodoShow)
	s.Router.HandleFunc("/todos", TodoIndex)
	s.Router.HandleFunc("/", Index)
}
