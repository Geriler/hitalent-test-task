package department

import (
	"errors"
	"fmt"

	"github.com/Geriler/hitalent/internal/domain/department"
	"github.com/Geriler/hitalent/internal/domain/employee"
)

type DeleteMode string

const (
	DeleteModeCascade  DeleteMode = "cascade"
	DeleteModeReassign DeleteMode = "reassign"
)

var (
	ErrRequireReassignToDepartmentID = errors.New("ReassignToDepartmentID is required")
)

type DeleteDepartmentUseCase struct {
	departmentRepo department.Repository
	employeeRepo   employee.Repository
}

type DeleteDepartmentUseCaseInput struct {
	ID                     int64
	Mode                   DeleteMode
	ReassignToDepartmentID *int64
}

type DeleteDepartmentUseCaseOutput struct{}

func NewDeleteDepartmentUseCase(departmentRepo department.Repository, employeeRepo employee.Repository) *DeleteDepartmentUseCase {
	return &DeleteDepartmentUseCase{
		departmentRepo: departmentRepo,
		employeeRepo:   employeeRepo,
	}
}

func (uc *DeleteDepartmentUseCase) Execute(input DeleteDepartmentUseCaseInput) (*DeleteDepartmentUseCaseOutput, error) {
	_, err := uc.departmentRepo.GetByID(input.ID)
	if err != nil {
		return nil, ErrNotFound
	}

	if input.Mode == DeleteModeReassign {
		if input.ReassignToDepartmentID == nil {
			return nil, ErrRequireReassignToDepartmentID
		}

		_, err = uc.departmentRepo.GetByID(*input.ReassignToDepartmentID)
		if err != nil {
			return nil, ErrNotFound
		}

		err = uc.employeeRepo.ReassignEmployees(input.ID, *input.ReassignToDepartmentID)
		if err != nil {
			return nil, fmt.Errorf("reassign employees: %w", err)
		}
	}

	err = uc.departmentRepo.Delete(input.ID)
	if err != nil {
		return nil, fmt.Errorf("delete department: %w", err)
	}

	return &DeleteDepartmentUseCaseOutput{}, nil
}
