package department

import (
	"testing"

	"github.com/Geriler/hitalent/internal/domain/department"
	mockdepartment "github.com/Geriler/hitalent/internal/mocks/department"
	mockemployee "github.com/Geriler/hitalent/internal/mocks/employee"
	"github.com/stretchr/testify/assert"
)

func TestDeleteDepartmentUseCase_Execute_ErrNotFound(t *testing.T) {
	deptRepo := mockdepartment.NewMockRepository(t)
	emplRepo := mockemployee.NewMockRepository(t)

	uc := NewDeleteDepartmentUseCase(deptRepo, emplRepo)

	id := int64(1)
	mode := DeleteMode("cascade")

	deptRepo.EXPECT().GetByID(id).Return(nil, ErrNotFound)

	result, err := uc.Execute(DeleteDepartmentUseCaseInput{
		ID:                     id,
		Mode:                   mode,
		ReassignToDepartmentID: nil,
	})

	assert.ErrorIs(t, err, ErrNotFound)
	assert.Nil(t, result)
}

func TestDeleteDepartmentUseCase_Execute_ErrRequireReassignToDepartmentID(t *testing.T) {
	deptRepo := mockdepartment.NewMockRepository(t)
	emplRepo := mockemployee.NewMockRepository(t)

	uc := NewDeleteDepartmentUseCase(deptRepo, emplRepo)

	id := int64(1)
	mode := DeleteMode("reassign")

	deptRepo.EXPECT().GetByID(id).Return(&department.Department{ID: 1, Name: "Test"}, nil)

	result, err := uc.Execute(DeleteDepartmentUseCaseInput{
		ID:                     id,
		Mode:                   mode,
		ReassignToDepartmentID: nil,
	})

	assert.ErrorIs(t, err, ErrRequireReassignToDepartmentID)
	assert.Nil(t, result)
}

func TestDeleteDepartmentUseCase_Execute_ErrNotFound_ReassignToDepartmentID(t *testing.T) {
	deptRepo := mockdepartment.NewMockRepository(t)
	emplRepo := mockemployee.NewMockRepository(t)

	uc := NewDeleteDepartmentUseCase(deptRepo, emplRepo)

	id := int64(1)
	mode := DeleteMode("reassign")
	reassignID := int64(2)

	deptRepo.EXPECT().GetByID(id).Return(&department.Department{ID: 1, Name: "Test 1"}, nil)
	deptRepo.EXPECT().GetByID(reassignID).Return(nil, ErrNotFound)

	result, err := uc.Execute(DeleteDepartmentUseCaseInput{
		ID:                     id,
		Mode:                   mode,
		ReassignToDepartmentID: &reassignID,
	})

	assert.ErrorIs(t, err, ErrNotFound)
	assert.Nil(t, result)
}

func TestDeleteDepartmentUseCase_Execute_Success_Cascade(t *testing.T) {
	deptRepo := mockdepartment.NewMockRepository(t)
	emplRepo := mockemployee.NewMockRepository(t)

	uc := NewDeleteDepartmentUseCase(deptRepo, emplRepo)

	id := int64(1)
	mode := DeleteMode("cascade")

	deptRepo.EXPECT().GetByID(id).Return(&department.Department{ID: 1, Name: "Test"}, nil)
	deptRepo.EXPECT().Delete(id).Return(nil)

	result, err := uc.Execute(DeleteDepartmentUseCaseInput{
		ID:                     id,
		Mode:                   mode,
		ReassignToDepartmentID: nil,
	})

	assert.NoError(t, err)
	assert.Equal(t, &DeleteDepartmentUseCaseOutput{}, result)
}

func TestDeleteDepartmentUseCase_Execute_Success_Reassign(t *testing.T) {
	deptRepo := mockdepartment.NewMockRepository(t)
	emplRepo := mockemployee.NewMockRepository(t)

	uc := NewDeleteDepartmentUseCase(deptRepo, emplRepo)

	id := int64(1)
	mode := DeleteMode("reassign")
	reassignID := int64(2)

	deptRepo.EXPECT().GetByID(id).Return(&department.Department{ID: 1, Name: "Test 1"}, nil)
	deptRepo.EXPECT().GetByID(reassignID).Return(&department.Department{ID: 2, Name: "Test 2"}, nil)
	emplRepo.EXPECT().ReassignEmployees(id, reassignID).Return(nil)
	deptRepo.EXPECT().Delete(id).Return(nil)

	result, err := uc.Execute(DeleteDepartmentUseCaseInput{
		ID:                     id,
		Mode:                   mode,
		ReassignToDepartmentID: &reassignID,
	})

	assert.NoError(t, err)
	assert.Equal(t, &DeleteDepartmentUseCaseOutput{}, result)
}
