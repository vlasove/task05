package apiserver

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vlasove/test05/internal/app/store"
)

// APIServer ...
type APIServer struct {
	config *Config
	router *mux.Router
	store  *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	log.Println("starting api server at port", s.config.BindAddr)

	s.configureRouter()
	log.Println("router configurated successfully")

	log.Println("starting connection to database...")
	if err := s.configureStore(); err != nil {
		return err
	}
	log.Println("database successfully connected")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// configureStore ...
func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

// configureRouter ...
func (a *APIServer) configureRouter() {
	a.router.HandleFunc("/hello", a.handleHello())
}

// handleHello ...
func (a *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
