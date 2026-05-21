package department

import (
	"testing"

	"github.com/Geriler/hitalent/internal/domain/department"
	"github.com/Geriler/hitalent/internal/domain/employee"
	mockdepartment "github.com/Geriler/hitalent/internal/mocks/department"
	mockemployee "github.com/Geriler/hitalent/internal/mocks/employee"
	"github.com/stretchr/testify/assert"
)

func TestGetDepartmentUseCase_Execute_ErrNotFound(t *testing.T) {
	deptRepo := mockdepartment.NewMockRepository(t)
	emplRepo := mockemployee.NewMockRepository(t)

	uc := NewGetDepartmentUseCase(deptRepo, emplRepo)

	id := int64(1)

	deptRepo.EXPECT().GetByID(id).Return(nil, ErrNotFound)

	result, err := uc.Execute(GetDepartmentUseCaseInput{
		DepartmentID:     id,
		Depth:            1,
		IncludeEmployees: false,
	})

	assert.ErrorIs(t, err, ErrNotFound)
	assert.Nil(t, result)
}

func TestGetDepartmentUseCase_Execute_Success(t *testing.T) {
	deptRepo := mockdepartment.NewMockRepository(t)
	emplRepo := mockemployee.NewMockRepository(t)

	uc := NewGetDepartmentUseCase(deptRepo, emplRepo)

	id := int64(1)
	depth := 1

	expectedDept := &department.Department{ID: id, Name: "Test"}

	deptRepo.EXPECT().GetByID(id).Return(expectedDept, nil)
	deptRepo.EXPECT().GetChildren(id, depth).Return([]*department.Department{}, nil)

	result, err := uc.Execute(GetDepartmentUseCaseInput{
		DepartmentID:     id,
		Depth:            depth,
		IncludeEmployees: false,
	})

	assert.NoError(t, err)
	assert.Equal(t, &GetDepartmentUseCaseOutput{
		Department: expectedDept,
		Children:   []*department.Department{},
		Employees:  nil,
	}, result)
}

func TestGetDepartmentUseCase_Execute_Success_IncludeEmployees(t *testing.T) {
	deptRepo := mockdepartment.NewMockRepository(t)
	emplRepo := mockemployee.NewMockRepository(t)

	uc := NewGetDepartmentUseCase(deptRepo, emplRepo)

	id := int64(1)
	depth := 1

	expectedDept := &department.Department{ID: id, Name: "Test"}

	deptRepo.EXPECT().GetByID(id).Return(expectedDept, nil)
	deptRepo.EXPECT().GetChildren(id, depth).Return([]*department.Department{}, nil)
	emplRepo.EXPECT().GetByDepartmentID(id).Return([]*employee.Employee{}, nil)

	result, err := uc.Execute(GetDepartmentUseCaseInput{
		DepartmentID:     id,
		Depth:            depth,
		IncludeEmployees: true,
	})

	assert.NoError(t, err)
	assert.Equal(t, &GetDepartmentUseCaseOutput{
		Department: expectedDept,
		Children:   []*department.Department{},
		Employees:  []*employee.Employee{},
	}, result)
}
