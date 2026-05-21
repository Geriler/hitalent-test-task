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

func TestEmployeeHandler_Create(t *testing.T) {
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

	fullName := "Alex"
	position := "Backend developer"

	body = fmt.Sprintf("{\"full_name\": \"%s\", \"position\": \"%s\"}", fullName, position)

	req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("/departments/%d/employees", respCreate.Department.ID), strings.NewReader(body))
	if err != nil {
		t.Errorf("unable to create request: %v", err)
	}
	w = httptest.NewRecorder()

	router.Routes().ServeHTTP(w, req)

	t.Log(w.Body.String())

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp handler.CreateEmployeeResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, resp.Employee.DepartmentID, respCreate.Department.ID)
	assert.Equal(t, resp.Employee.FullName, fullName)
	assert.Equal(t, resp.Employee.Position, position)
}
