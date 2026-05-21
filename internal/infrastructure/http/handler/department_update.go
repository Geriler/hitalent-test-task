package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	departmentUseCases "github.com/Geriler/hitalent/internal/usecase/department"
)

type UpdateDepartmentRequest struct {
	ID       int64   `json:"-"`
	Name     *string `json:"name"`
	ParentID *int64  `json:"parent_id"`
}

type UpdateDepartmentResponse struct {
	Department *DepartmentDTO `json:"department"`
}

func (h *DepartmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	req := UpdateDepartmentRequest{
		ID: id,
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.updateUC.Execute(departmentUseCases.UpdateDepartmentUseCaseInput{
		ID:       req.ID,
		Name:     req.Name,
		ParentID: req.ParentID,
	})
	if err != nil {
		httpError(w, err)
		return
	}

	resp := UpdateDepartmentResponse{
		Department: toDepartmentDTO(output.Department),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
