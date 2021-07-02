package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	public := s.router.PathPrefix("/api/v1").Subrouter()
	//public.Use(s.checkContentType)
	public.Use(s.setContentType)
	public.HandleFunc("/employees", s.handleRetrieveAll()).Methods("GET")
	public.HandleFunc("/employees/{employeeId:[0-9]+}", s.handleRetrieveByID()).Methods("GET")
	public.HandleFunc("/employees", s.handleCreate()).Methods("POST")
	public.HandleFunc("/employees/{employeeId:[0-9]+}", s.handleDelete()).Methods("DELETE")
	public.HandleFunc("/employees/{employeeId:[0-9]+}", s.handleUpdate()).Methods("PUT")

	technical := s.router.PathPrefix("/tech").Subrouter()
	technical.HandleFunc("/info", s.handleInfo()).Methods("GET")
}

func (s *server) handleUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["employeeId"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		e, err := s.store.Employee().GetByID(r.Context(), id)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		err = s.store.Employee().Update(r.Context(), e)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusAccepted, map[string]string{"message": "employee's info updated"})
	}
}

func (s *server) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["employeeId"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := s.store.Employee().Delete(r.Context(), id); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusAccepted, map[string]string{"message": "emplyee deleted"})
	}
}

func (s *server) handleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := new(model.Employee)
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := s.store.Employee().Create(r.Context(), e); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		s.respond(w, r, http.StatusCreated, map[string]string{"message": "employee created"})
	}
}

func (s *server) handleRetrieveByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["employeeId"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		empl, err := s.store.Employee().GetByID(r.Context(), id)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, empl)
	}
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
