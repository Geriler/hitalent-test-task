package department

import (
	"errors"
	"fmt"

	"github.com/Geriler/hitalent/internal/domain/department"
)

var (
	ErrParentNotFound = errors.New("parent department not found")
	ErrAlreadyExists  = errors.New("department already exists")
)

type CreateDepartmentUseCase struct {
	repo department.Repository
}

type CreateDepartmentUseCaseInput struct {
	Name     string
	ParentID *int64
}

type CreateDepartmentUseCaseOutput struct {
	Department *department.Department
}

func NewCreateDepartmentUseCase(repo department.Repository) *CreateDepartmentUseCase {
	return &CreateDepartmentUseCase{repo: repo}
}

func (uc *CreateDepartmentUseCase) Execute(input CreateDepartmentUseCaseInput) (*CreateDepartmentUseCaseOutput, error) {
	dept, err := department.NewDepartment(input.Name, input.ParentID)
	if err != nil {
		return nil, err
	}

	if input.ParentID != nil {
		_, err = uc.repo.GetByID(*input.ParentID)
		if err != nil {
			return nil, ErrParentNotFound
		}
	}

	exist, err := uc.repo.ExistsByNameAndParent(input.Name, input.ParentID)
	if err != nil {
		return nil, fmt.Errorf("check department existence: %w", err)
	}
	if exist {
		return nil, ErrAlreadyExists
	}

	resp, err := uc.repo.Create(dept)
	if err != nil {
		return nil, fmt.Errorf("create department: %w", err)
	}

	return &CreateDepartmentUseCaseOutput{
		Department: resp,
	}, nil
}
