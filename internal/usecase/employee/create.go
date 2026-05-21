package employee

import (
	"errors"
	"fmt"
	"time"

	"github.com/Geriler/hitalent/internal/domain/department"
	"github.com/Geriler/hitalent/internal/domain/employee"
)

var ErrDepartmentNotFound = errors.New("department not found")

type CreateEmployeeUseCase struct {
	employeeRepo   employee.Repository
	departmentRepo department.Repository
}

type CreateEmployeeUseCaseInput struct {
	DepartmentID int64
	FullName     string
	Position     string
	HiredAt      *time.Time
}

type CreateEmployeeUseCaseOutput struct {
	Employee *employee.Employee
}

func NewCreateEmployeeUseCase(employeeRepo employee.Repository, departmentRepo department.Repository) *CreateEmployeeUseCase {
	return &CreateEmployeeUseCase{
		employeeRepo:   employeeRepo,
		departmentRepo: departmentRepo,
	}
}

func (uc *CreateEmployeeUseCase) Execute(input CreateEmployeeUseCaseInput) (*CreateEmployeeUseCaseOutput, error) {
	_, err := uc.departmentRepo.GetByID(input.DepartmentID)
	if err != nil {
		return nil, ErrDepartmentNotFound
	}

	empl, err := employee.NewEmployee(input.DepartmentID, input.FullName, input.Position, input.HiredAt)
	if err != nil {
		return nil, err
	}

	resp, err := uc.employeeRepo.Create(empl)
	if err != nil {
		return nil, fmt.Errorf("create employee: %w", err)
	}

	return &CreateEmployeeUseCaseOutput{Employee: resp}, nil
}
