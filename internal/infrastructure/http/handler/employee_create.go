package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Geriler/hitalent/internal/usecase/employee"
)

type CreateEmployeeRequest struct {
	FullName string     `json:"full_name"`
	Position string     `json:"position"`
	HiredAt  *time.Time `json:"hired_at"`
}

type CreateEmployeeResponse struct {
	Employee *EmployeeDTO `json:"employee"`
}

func (h *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	deptIDStr := r.PathValue("id")
	deptID, err := strconv.ParseInt(deptIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	req := &CreateEmployeeRequest{}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.createUC.Execute(employee.CreateEmployeeUseCaseInput{
		DepartmentID: deptID,
		FullName:     req.FullName,
		Position:     req.Position,
		HiredAt:      req.HiredAt,
	})
	if err != nil {
		httpError(w, err)
		return
	}

	resp := &CreateEmployeeResponse{
		Employee: toEmployeeDTO(output.Employee),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
