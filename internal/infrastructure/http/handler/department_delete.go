package handler

import (
	"net/http"
	"strconv"

	"github.com/Geriler/hitalent/internal/usecase/department"
)

func (h *DepartmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	mode := department.DeleteMode("cascade")
	if m := r.URL.Query().Get("mode"); m != "" {
		mode = department.DeleteMode(m)
		if mode != department.DeleteModeCascade && mode != department.DeleteModeReassign {
			http.Error(w, "invalid mode, must be cascade or reassign", http.StatusBadRequest)
			return
		}
	}

	reassignID := (*int64)(nil)
	if rID := r.URL.Query().Get("reassign_to_department_id"); rID != "" {
		intReassignID, err := strconv.ParseInt(rID, 10, 64)
		if err != nil {
			http.Error(w, "invalid reassign id", http.StatusBadRequest)
			return
		}
		reassignID = &intReassignID
	}

	_, err = h.deleteUC.Execute(department.DeleteDepartmentUseCaseInput{
		ID:                     id,
		Mode:                   mode,
		ReassignToDepartmentID: reassignID,
	})
	if err != nil {
		httpError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
