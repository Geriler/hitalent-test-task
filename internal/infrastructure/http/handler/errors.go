package handler

import (
	"errors"
	"net/http"

	"github.com/Geriler/hitalent/internal/domain/department"
	"github.com/Geriler/hitalent/internal/domain/employee"
	departmentUseCases "github.com/Geriler/hitalent/internal/usecase/department"
)

func httpError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, departmentUseCases.ErrNotFound),
		errors.Is(err, departmentUseCases.ErrParentNotFound),
		errors.Is(err, department.ErrNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	case errors.Is(err, departmentUseCases.ErrAlreadyExists),
		errors.Is(err, departmentUseCases.ErrCyclicDependency),
		errors.Is(err, departmentUseCases.ErrSelfParent):
		http.Error(w, err.Error(), http.StatusConflict)
	case errors.Is(err, departmentUseCases.ErrRequireReassignToDepartmentID),
		errors.Is(err, employee.ErrInvalidFullName),
		errors.Is(err, employee.ErrInvalidPosition),
		errors.Is(err, department.ErrInvalidName):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
