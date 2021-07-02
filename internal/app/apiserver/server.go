package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vlasove/test05/internal/app/store"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}
	s.configureRouter()
	return s
}

// ServerHTTP ...
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	public := s.router.PathPrefix("/api/v1").Subrouter()
	public.HandleFunc("/hello", s.handleHello()).Methods("GET")

	technical := s.router.PathPrefix("/tech").Subrouter()
	technical.HandleFunc("/info", s.handleInfo()).Methods("GET")
}

// handleHello ...
func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "hello")
	}
}

func (s *server) handleInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
