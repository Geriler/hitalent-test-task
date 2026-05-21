package department

import (
	"fmt"

	"github.com/Geriler/hitalent/internal/domain/department"
	"github.com/Geriler/hitalent/internal/domain/employee"
)

type GetDepartmentUseCase struct {
	departmentRepo department.Repository
	employeeRepo   employee.Repository
}

type GetDepartmentUseCaseInput struct {
	DepartmentID     int64
	Depth            int
	IncludeEmployees bool
}

type GetDepartmentUseCaseOutput struct {
	Department *department.Department
	Children   []*department.Department
	Employees  []*employee.Employee
}

func NewGetDepartmentUseCase(departmentRepo department.Repository, employeeRepo employee.Repository) *GetDepartmentUseCase {
	return &GetDepartmentUseCase{
		departmentRepo: departmentRepo,
		employeeRepo:   employeeRepo,
	}
}

func (uc *GetDepartmentUseCase) Execute(input GetDepartmentUseCaseInput) (*GetDepartmentUseCaseOutput, error) {
	dept, err := uc.departmentRepo.GetByID(input.DepartmentID)
	if err != nil {
		return nil, fmt.Errorf("get department by id: %w", err)
	}

	children, err := uc.departmentRepo.GetChildren(input.DepartmentID, input.Depth)
	if err != nil {
		return nil, fmt.Errorf("get department children: %w", err)
	}

	var employees []*employee.Employee
	if input.IncludeEmployees {
		employees, err = uc.employeeRepo.GetByDepartmentID(dept.ID)
		if err != nil {
			return nil, fmt.Errorf("get employees by department id: %w", err)
		}
	}

	return &GetDepartmentUseCaseOutput{
		Department: dept,
		Children:   children,
		Employees:  employees,
	}, nil
}
