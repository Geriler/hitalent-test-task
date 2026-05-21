package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Geriler/hitalent/internal/infrastructure/http/handler"
	"github.com/stretchr/testify/assert"
)

func TestDepartmentHandler_Create(t *testing.T) {
	router := setupRouter(t)

	body := `{"name": "Backend"}`

	req, err := http.NewRequest(http.MethodPost, "/departments", strings.NewReader(body))
	if err != nil {
		t.Errorf("unable to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.Routes().ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp handler.CreateDepartmentResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Backend", resp.Department.Name)
	assert.Nil(t, resp.Department.ParentID)
}

func TestDepartmentHandler_Get(t *testing.T) {
	router := setupRouter(t)

	body := `{"name": "Backend"}`

	req, err := http.NewRequest(http.MethodPost, "/departments", strings.NewReader(body))
	if err != nil {
		t.Errorf("unable to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.Routes().ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var respCreate handler.CreateDepartmentResponse
	json.Unmarshal(w.Body.Bytes(), &respCreate)

	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("/departments/%d", respCreate.Department.ID), nil)
	if err != nil {
		t.Errorf("unable to create request: %v", err)
	}
	w = httptest.NewRecorder()

	router.Routes().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp handler.GetDepartmentResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Backend", resp.Department.Name)
	assert.Nil(t, resp.Department.ParentID)
}

func TestDepartmentHandler_Update(t *testing.T) {
	router := setupRouter(t)

	body := `{"name": "Backend"}`

	req, err := http.NewRequest(http.MethodPost, "/departments", strings.NewReader(body))
	if err != nil {
		t.Errorf("unable to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.Routes().ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var respCreate handler.CreateDepartmentResponse
	json.Unmarshal(w.Body.Bytes(), &respCreate)

	body = `{"name": "QA"}`

	req, err = http.NewRequest(http.MethodPatch, fmt.Sprintf("/departments/%d", respCreate.Department.ID), strings.NewReader(body))
	if err != nil {
		t.Errorf("unable to create request: %v", err)
	}
	w = httptest.NewRecorder()

	router.Routes().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp handler.UpdateDepartmentResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "QA", resp.Department.Name)
	assert.Nil(t, resp.Department.ParentID)
}

func TestDepartmentHandler_Delete(t *testing.T) {
	router := setupRouter(t)

	body := `{"name": "Backend"}`

	req, err := http.NewRequest(http.MethodPost, "/departments", strings.NewReader(body))
	if err != nil {
		t.Errorf("unable to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.Routes().ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var respCreate handler.CreateDepartmentResponse
	json.Unmarshal(w.Body.Bytes(), &respCreate)

	req, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("/departments/%d", respCreate.Department.ID), nil)
	if err != nil {
		t.Errorf("unable to create request: %v", err)
	}
	w = httptest.NewRecorder()

	router.Routes().ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("/departments/%d", respCreate.Department.ID), nil)
	if err != nil {
		t.Errorf("unable to create request: %v", err)
	}
	w = httptest.NewRecorder()

	router.Routes().ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
