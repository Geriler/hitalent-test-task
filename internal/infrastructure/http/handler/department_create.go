package handler

import (
	"encoding/json"
	"net/http"

	departmentUseCases "github.com/Geriler/hitalent/internal/usecase/department"
)

type CreateDepartmentRequest struct {
	Name     string `json:"name"`
	ParentID *int64 `json:"parent_id"`
}

type CreateDepartmentResponse struct {
	Department *DepartmentDTO `json:"department"`
}

func (h *DepartmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := CreateDepartmentRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.createUC.Execute(departmentUseCases.CreateDepartmentUseCaseInput{
		Name:     req.Name,
		ParentID: req.ParentID,
	})
	if err != nil {
		httpError(w, err)
		return
	}

	resp := CreateDepartmentResponse{
		Department: toDepartmentDTO(output.Department),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
