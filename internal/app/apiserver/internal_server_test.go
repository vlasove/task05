package apiserver

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vlasove/test05/internal/app/model"
	"github.com/vlasove/test05/internal/app/store/teststore"
)

func TestServer_HandleInfo(t *testing.T) {
	store := teststore.New()
	s := newServer(store)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/tech/info", nil)
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)

	var v map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&v)
	assert.NoError(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, v["name"], techName)
	assert.Equal(t, v["version"], techVersion)

}

func TestServer_HandleRetrieveAll(t *testing.T) {
	t.Run("empty store", func(t *testing.T) {
		store := teststore.New()
		s := newServer(store)
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		s.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, http.StatusOK)

		var v map[string][]*model.Employee
		err := json.NewDecoder(rec.Body).Decode(&v)
		assert.NoError(t, err)
		assert.NotNil(t, v)
		assert.Equal(t, len(v["employees"]), 0)

	})

	t.Run("non empty store", func(t *testing.T) {
		store := teststore.New()
		_ = store.Employee().Create(context.Background(), model.TestEmployee(t))
		_ = store.Employee().Create(context.Background(), model.TestEmployee(t))
		s := newServer(store)
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		s.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, http.StatusOK)

		var v map[string][]*model.Employee
		err := json.NewDecoder(rec.Body).Decode(&v)
		assert.NoError(t, err)
		assert.NotNil(t, v)
		assert.Equal(t, len(v["employees"]), 2)
	})
}
