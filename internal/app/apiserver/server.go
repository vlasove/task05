package apiserver

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vlasove/test05/internal/app/model"
	"github.com/vlasove/test05/internal/app/store"
)

const (
	techName    = "employees"
	techVersion = "1.0.0"
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
	s.router.Use(s.logRequest)
	s.router.Use(s.setContentType)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	public := s.router.PathPrefix("/api/v1").Subrouter()
	public.HandleFunc("/hello", s.handleHello()).Methods("GET")
	public.HandleFunc("/employees", s.handleRetrieveAll()).Methods("GET")

	technical := s.router.PathPrefix("/tech").Subrouter()
	technical.HandleFunc("/info", s.handleInfo()).Methods("GET")
}

func (s *server) handleRetrieveAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		empls, err := s.store.Employee().GetAll(r.Context())
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, map[string][]*model.Employee{"employees": empls})
	}
}

// handleHello ...
func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "hello")
	}
}

// handleInfo ...
func (s *server) handleInfo() http.HandlerFunc {
	type response struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	return func(w http.ResponseWriter, r *http.Request) {

		s.respond(w, r, http.StatusOK, response{
			Name:    techName,
			Version: techVersion,
		})
	}
}

// respond ...
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}

}

// error ...
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}
