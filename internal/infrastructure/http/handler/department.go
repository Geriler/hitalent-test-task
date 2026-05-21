package handler

import (
	departmentUseCases "github.com/Geriler/hitalent/internal/usecase/department"
)

type DepartmentHandler struct {
	createUC *departmentUseCases.CreateDepartmentUseCase
	getUC    *departmentUseCases.GetDepartmentUseCase
	updateUC *departmentUseCases.UpdateDepartmentUseCase
	deleteUC *departmentUseCases.DeleteDepartmentUseCase
}

func NewDepartmentHandler(
	createUC *departmentUseCases.CreateDepartmentUseCase,
	getUC *departmentUseCases.GetDepartmentUseCase,
	updateUC *departmentUseCases.UpdateDepartmentUseCase,
	deleteUC *departmentUseCases.DeleteDepartmentUseCase,
) *DepartmentHandler {
	return &DepartmentHandler{
		createUC: createUC,
		getUC:    getUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
	}
}
