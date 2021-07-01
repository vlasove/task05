package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_handleHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "hello")
}
