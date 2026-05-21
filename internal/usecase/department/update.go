package department

import (
	"errors"
	"fmt"
	"slices"

	"github.com/Geriler/hitalent/internal/domain/department"
)

var (
	ErrSelfParent       = errors.New("department cannot be its own parent")
	ErrCyclicDependency = errors.New("cannot move department into its own subtree")
	ErrNotFound         = errors.New("department not found")
)

type UpdateDepartmentUseCase struct {
	repo department.Repository
}

type UpdateDepartmentUseCaseInput struct {
	ID       int64
	Name     *string
	ParentID *int64
}

type UpdateDepartmentUseCaseOutput struct {
	Department *department.Department
}

func NewUpdateDepartmentUseCase(repo department.Repository) *UpdateDepartmentUseCase {
	return &UpdateDepartmentUseCase{repo: repo}
}

func (uc *UpdateDepartmentUseCase) Execute(input UpdateDepartmentUseCaseInput) (*UpdateDepartmentUseCaseOutput, error) {
	dept, err := uc.repo.GetByID(input.ID)
	if err != nil {
		return nil, ErrNotFound
	}

	err = dept.Update(input.Name, input.ParentID)
	if err != nil {
		return nil, err
	}

	if input.ParentID != nil {
		if *input.ParentID == input.ID {
			return nil, ErrSelfParent
		}

		subtreeIDs, err := uc.repo.GetSubtreeIDs(input.ID)
		if err != nil {
			return nil, fmt.Errorf("get subtree ids: %w", err)
		}

		if slices.Contains(subtreeIDs, *input.ParentID) {
			return nil, ErrCyclicDependency
		}
	}

	resp, err := uc.repo.Update(input.ID, dept)
	if err != nil {
		return nil, fmt.Errorf("update department: %w", err)
	}

	return &UpdateDepartmentUseCaseOutput{
		Department: resp,
	}, nil
}
