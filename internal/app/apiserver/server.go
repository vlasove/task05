package apiserver

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
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
	s.router.Use(s.setContentLang)
	s.router.Use(s.setContentType)
	s.router.Use(s.authMiddleware)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	public := s.router.PathPrefix("/api/v1").Subrouter()

	public.HandleFunc("/employees/{employeeId:[0-9]+}", s.handleRetrieveByID()).Methods("GET")
	public.HandleFunc("/employees/{employeeId:[0-9]+}", s.handleDelete()).Methods("DELETE")

	acceptbl := s.router.PathPrefix("/api/v1").Subrouter()
	acceptbl.Use(s.acceptMiddleware)
	acceptbl.HandleFunc("/employees", s.handleRetrieveAll()).Methods("GET")

	jsonable := s.router.PathPrefix("/api/v1").Subrouter()
	jsonable.Use(s.checkContentType)
	jsonable.HandleFunc("/employees", s.handleCreate()).Methods("POST")
	jsonable.HandleFunc("/employees/{employeeId:[0-9]+}", s.handleUpdate()).Methods("PUT")

	technical := s.router.PathPrefix("/tech").Subrouter()
	technical.HandleFunc("/info", s.handleInfo()).Methods("GET")
}

func (s *server) handleUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["employeeId"])
		if err != nil {
			s.error(w, r, err)
			return
		}

		e, err := s.store.Employee().GetByID(r.Context(), id)
		if err != nil {
			s.error(w, r, err)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			s.error(w, r, err)
			return
		}

		err = s.store.Employee().Update(r.Context(), e)
		if err != nil {
			s.error(w, r, err)
			return
		}

		s.respond(w, r, http.StatusAccepted, map[string]string{"message": "employee's info updated"})
	}
}

func (s *server) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["employeeId"])
		if err != nil {
			s.error(w, r, err)
			return
		}
		if err := s.store.Employee().Delete(r.Context(), id); err != nil {
			s.error(w, r, err)
			return
		}
		s.respond(w, r, http.StatusAccepted, map[string]string{"message": "emplyee deleted"})
	}
}

func (s *server) handleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := new(model.Employee)
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			s.error(w, r, err)
			return
		}
		if err := s.store.Employee().Create(r.Context(), e); err != nil {
			s.error(w, r, err)
			return
		}
		s.respond(w, r, http.StatusCreated, map[string]string{"message": "employee created"})
	}
}

func (s *server) handleRetrieveByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["employeeId"])
		if err != nil {
			s.error(w, r, err)
			return
		}
		empl, err := s.store.Employee().GetByID(r.Context(), id)
		if err != nil {
			s.error(w, r, err)
			return
		}
		s.respond(w, r, http.StatusOK, empl)
	}
}

func (s *server) handleRetrieveAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		empls, err := s.store.Employee().GetAll(r.Context())
		if err != nil {
			s.error(w, r, err)
			return
		}

		s.respond(w, r, http.StatusOK, empls)

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
		switch {
		case r.Context().Value(ctxXMLKey):
			_ = xml.NewEncoder(w).Encode(data)
		default:
			_ = json.NewEncoder(w).Encode(data)
		}

	}

}

// error ...
func (s *server) error(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/problem+json")
	problemDescription := make(map[string]interface{})
	switch err := err.(type) {
	case *pq.Error:
		if err.Code >= "50000" {
			problemDescription["title"] = "client error"
			problemDescription["detail"] = err.Message
			problemDescription["instance"] = r.URL.String()
			s.respond(w, r, http.StatusBadRequest, problemDescription)
			return
		}
		problemDescription["title"] = "internal error"
		problemDescription["detail"] = err.Message
		problemDescription["instance"] = r.URL.String()
		s.respond(w, r, http.StatusInternalServerError, problemDescription)
		return

	case validation.Errors:
		problemDescription["title"] = "validation error"
		problemDescription["detail"] = "you used unapropriate content in your JSON payload data"
		problemDescription["instance"] = r.URL.String()

		temp := make([]map[string]string, 0)
		for k, v := range err {
			temp = append(temp,
				map[string]string{k: v.Error()},
			)
		}
		problemDescription["fields"] = temp

		s.respond(w, r, http.StatusBadRequest, problemDescription)
		return

	default:
		problemDescription["title"] = "invalid input"
		problemDescription["detail"] = "unacceptble URL params or invalid JSON body:" + err.Error()
		problemDescription["instance"] = r.URL.String()
		s.respond(w, r, http.StatusBadRequest, problemDescription)
		return

	}
}
