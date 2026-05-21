package employee

import (
	"testing"

	"github.com/Geriler/hitalent/internal/domain/department"
	"github.com/Geriler/hitalent/internal/domain/employee"
	mockdepartment "github.com/Geriler/hitalent/internal/mocks/department"
	mockemployee "github.com/Geriler/hitalent/internal/mocks/employee"
	"github.com/stretchr/testify/assert"
)

func TestCreateEmployeeUseCase_Execute_ErrDepartmentNotFound(t *testing.T) {
	deptRepo := mockdepartment.NewMockRepository(t)
	emplRepo := mockemployee.NewMockRepository(t)

	uc := NewCreateEmployeeUseCase(emplRepo, deptRepo)

	departmentID := int64(1)

	deptRepo.EXPECT().GetByID(departmentID).Return(nil, ErrDepartmentNotFound)

	result, err := uc.Execute(CreateEmployeeUseCaseInput{
		DepartmentID: departmentID,
	})

	assert.ErrorIs(t, err, ErrDepartmentNotFound)
	assert.Nil(t, result)
}

func TestCreateEmployeeUseCase_Execute_ErrInvalidFullName(t *testing.T) {
	deptRepo := mockdepartment.NewMockRepository(t)
	emplRepo := mockemployee.NewMockRepository(t)

	uc := NewCreateEmployeeUseCase(emplRepo, deptRepo)

	departmentID := int64(1)
	fullName := ""

	deptRepo.EXPECT().GetByID(departmentID).Return(&department.Department{ID: departmentID, Name: "Test"}, nil)

	result, err := uc.Execute(CreateEmployeeUseCaseInput{
		DepartmentID: departmentID,
		FullName:     fullName,
	})

	assert.ErrorIs(t, err, employee.ErrInvalidFullName)
	assert.Nil(t, result)
}

func TestCreateEmployeeUseCase_Execute_ErrInvalidPosition(t *testing.T) {
	deptRepo := mockdepartment.NewMockRepository(t)
	emplRepo := mockemployee.NewMockRepository(t)

	uc := NewCreateEmployeeUseCase(emplRepo, deptRepo)

	departmentID := int64(1)
	fullName := "test"
	position := ""

	deptRepo.EXPECT().GetByID(departmentID).Return(&department.Department{ID: departmentID, Name: "Test"}, nil)

	result, err := uc.Execute(CreateEmployeeUseCaseInput{
		DepartmentID: departmentID,
		FullName:     fullName,
		Position:     position,
	})

	assert.ErrorIs(t, err, employee.ErrInvalidPosition)
	assert.Nil(t, result)
}

func TestCreateEmployeeUseCase_Execute_Success(t *testing.T) {
	deptRepo := mockdepartment.NewMockRepository(t)
	emplRepo := mockemployee.NewMockRepository(t)

	uc := NewCreateEmployeeUseCase(emplRepo, deptRepo)

	departmentID := int64(1)
	employeeID := int64(1)
	fullName := "test"
	position := "abc"
	expectedEmployee := &employee.Employee{
		ID:           employeeID,
		DepartmentID: departmentID,
		FullName:     fullName,
		Position:     position,
	}

	deptRepo.EXPECT().GetByID(departmentID).Return(&department.Department{ID: departmentID, Name: "Test"}, nil)
	emplRepo.EXPECT().Create(&employee.Employee{DepartmentID: departmentID, FullName: fullName, Position: position}).
		Return(expectedEmployee, nil)

	result, err := uc.Execute(CreateEmployeeUseCaseInput{
		DepartmentID: departmentID,
		FullName:     fullName,
		Position:     position,
	})

	assert.NoError(t, err)
	assert.Equal(t, &CreateEmployeeUseCaseOutput{Employee: expectedEmployee}, result)
}
