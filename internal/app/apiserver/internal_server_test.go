package apiserver

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
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
		s.handleRetrieveAll().ServeHTTP(rec, req)
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
		s.handleRetrieveAll().ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, http.StatusOK)

		var v map[string][]*model.Employee
		err := json.NewDecoder(rec.Body).Decode(&v)
		assert.NoError(t, err)
		assert.NotNil(t, v)
		assert.Equal(t, len(v["employees"]), 2)
	})
}

func TestServer_HandleRetrieveByID(t *testing.T) {

	t.Run("not existing id", func(t *testing.T) {
		store := teststore.New()
		s := newServer(store)
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		req = mux.SetURLVars(req, map[string]string{"employeeId": "1"})

		s.handleRetrieveByID().ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})

	t.Run("existing user", func(t *testing.T) {
		store := teststore.New()
		err := store.Employee().Create(context.Background(), model.TestEmployee(t))
		assert.NoError(t, err)
		s := newServer(store)
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		req = mux.SetURLVars(req, map[string]string{"employeeId": "1"})

		s.handleRetrieveByID().ServeHTTP(rec, req)

		assert.Equal(t, rec.Code, http.StatusOK)
		var v *model.Employee
		err = json.NewDecoder(rec.Body).Decode(&v)
		assert.NoError(t, err)
		assert.NotNil(t, v)

	})
}

func TestServer_HandleCreate(t *testing.T) {
	store := teststore.New()
	s := newServer(store)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"name":           "admin",
				"last_name":      "adminovich",
				"patronymic":     "patros",
				"phone":          "88005553535",
				"position":       "admin",
				"good_job_count": 10,
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid",
			payload: map[string]interface{}{
				"name":           "admin",
				"patronymic":     "patros",
				"phone":          "88005553535",
				"position":       "admin",
				"good_job_count": 10,
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			_ = json.NewEncoder(b).Encode(tt.payload)

			req, _ := http.NewRequest(http.MethodPost, "/api/v1/employees", b)
			s.handleCreate().ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tt.expectedCode)
		})
	}
}

func TestServer_HandleUpdate(t *testing.T) {
	store := teststore.New()
	err := store.Employee().Create(context.Background(), model.TestEmployee(t))
	assert.NoError(t, err)
	s := newServer(store)

	rec := httptest.NewRecorder()
	e := model.TestEmployee(t)
	e.ID = 1
	e.Name = "upd name"
	b := &bytes.Buffer{}
	_ = json.NewEncoder(b).Encode(e)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/employees", b)
	req = mux.SetURLVars(req, map[string]string{"employeeId": "1"})

	s.handleUpdate().ServeHTTP(rec, req)
	log.Println(rec)

	assert.Equal(t, rec.Code, http.StatusAccepted)

}

func TestServer_HandleDelete(t *testing.T) {
	store := teststore.New()
	err := store.Employee().Create(context.Background(), model.TestEmployee(t))
	assert.NoError(t, err)
	s := newServer(store)
	testCases := []struct {
		name         string
		id           string
		expectedCode int
	}{
		{
			name:         "valid",
			id:           "1",
			expectedCode: http.StatusAccepted,
		},
		{
			name:         "invalid",
			id:           "10",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, "/api/v1/employees", nil)
			req = mux.SetURLVars(req, map[string]string{"employeeId": tt.id})

			s.handleDelete().ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tt.expectedCode)

		})
	}
}
