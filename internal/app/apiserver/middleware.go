package apiserver

import (
	"errors"
	"log"
	"net/http"
	"time"
)

var (
	errContentTypeIsNotProvided = errors.New("content-type not provided")
	errInvalidContentType       = errors.New("use only content type : application/json")
)

// logRequest ...
func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("started %s %s\n", r.Method, r.RequestURI)
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)
		log.Printf("completed with %d %s in %v\n",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
		)

	})
}

// setContentType ...
func (s *server) setContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// setContentLang...
func (s *server) setContentLang(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Language", "ru-en")
		next.ServeHTTP(w, r)
	})
}

// checkContentType ...
// func (s *server) checkContentType(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodPost || r.Method == http.MethodPut {
// 			ct := w.Header().Get("Content-Type")
// 			log.Println("Content-type:", ct)
// 			switch {
// 			case len(ct) == 0:
// 				log.Println("trying to throw an error")
// 				s.error(w, r, http.StatusNoContent, errContentTypeIsNotProvided)
// 				return
// 			case ct != "application/json":
// 				s.error(w, r, http.StatusBadRequest, errInvalidContentType)
// 				return
// 			default:
// 				next.ServeHTTP(w, r)
// 				return
// 			}

// 		}
// 		next.ServeHTTP(w, r)
// 	})

// }
