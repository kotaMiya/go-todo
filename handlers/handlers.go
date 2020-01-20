package handlers

import "github.com/go-chi/chi"

type Sever struct {
}

func NewServer() *Server {
	return &Server{}
}

func SetupRouter() *chi.Mux {
	server := NewServer()

	r := chi.NewRouter()
	server.setupEndpoints(r)
	return r
}
