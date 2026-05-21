package handler

import "github.com/Geriler/hitalent/internal/usecase/employee"

type EmployeeHandler struct {
	createUC *employee.CreateEmployeeUseCase
}

func NewEmployeeHandler(createUC *employee.CreateEmployeeUseCase) *EmployeeHandler {
	return &EmployeeHandler{
		createUC: createUC,
	}
}
