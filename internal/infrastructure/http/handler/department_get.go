package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	departmentUseCases "github.com/Geriler/hitalent/internal/usecase/department"
)

type GetDepartmentResponse struct {
	Department *DepartmentDTO   `json:"department"`
	Children   []*DepartmentDTO `json:"children"`
	Employees  []*EmployeeDTO   `json:"employees"`
}

func (h *DepartmentHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	depth := 1
	if d := r.URL.Query().Get("depth"); d != "" {
		depth, err = strconv.Atoi(d)
		if err != nil || depth < 1 || depth > 5 {
			http.Error(w, "invalid depth", http.StatusBadRequest)
			return
		}
	}

	includeEmployees := false
	if ie := r.URL.Query().Get("include_employees"); ie != "" {
		includeEmployees, err = strconv.ParseBool(ie)
		if err != nil {
			http.Error(w, "invalid include_employees", http.StatusBadRequest)
			return
		}
	}

	output, err := h.getUC.Execute(departmentUseCases.GetDepartmentUseCaseInput{
		DepartmentID:     id,
		Depth:            depth,
		IncludeEmployees: includeEmployees,
	})
	if err != nil {
		httpError(w, err)
		return
	}

	resp := GetDepartmentResponse{
		Department: toDepartmentDTO(output.Department),
		Children:   toDepartmentDTOs(output.Children),
		Employees:  toEmployeeDTOs(output.Employees),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
