package apiserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type ctxKey string

const (
	ctxXMLKey ctxKey = "xml"
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
		w.Header().Set("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// setContentLang...
func (s *server) setContentLang(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-language", "ru-en")
		next.ServeHTTP(w, r)
	})
}

// checkContentType ...
func (s *server) checkContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("content-type") != "application/json" {

			s.respond(w, r, http.StatusUnsupportedMediaType,
				map[string]string{
					"title":    "invalid media type",
					"detail":   "use only application/json type for this request",
					"instance": r.URL.String(),
				},
			)

		} else {
			next.ServeHTTP(w, r)
		}

	})

}

// acceptMiddleware ...
func (s *server) acceptMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context
		if val := r.Header.Get("accept"); val == "application/xml" {
			w.Header().Set("content-type", "application/xml")
			ctx = context.WithValue(r.Context(), ctxXMLKey, true)
		} else {
			ctx = context.WithValue(r.Context(), ctxXMLKey, false)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// authMiddleware ...
func (s *server) authMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			s.respond(w, r, http.StatusUnauthorized,
				map[string]string{
					"title":    "invalid token",
					"detail":   "Bearer token invalid or not provided",
					"instance": r.URL.String(),
				},
			)
		} else {
			var secret []byte
			tokenString := authHeader[1]
			token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return secret, nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				val := reflect.ValueOf(claims["client_claims"])
				for i := 0; i < val.Len(); i++ {
					if val.Index(i).Interface().(string) == "world" {
						next.ServeHTTP(w, r)
						return
					}
				}
			} else {
				s.respond(w, r, http.StatusUnauthorized,
					map[string]string{
						"title":    "invalid token",
						"detail":   "Bearer token expired",
						"instance": r.URL.String(),
					},
				)
			}

			//next.ServeHTTP(w, r)
		}
	})
}
