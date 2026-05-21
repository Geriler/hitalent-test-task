package department

import (
	"testing"

	"github.com/Geriler/hitalent/internal/domain/department"
	mockdepartment "github.com/Geriler/hitalent/internal/mocks/department"
	"github.com/stretchr/testify/assert"
)

func TestUpdateDepartmentUseCase_Execute_ErrNotFound(t *testing.T) {
	repo := mockdepartment.NewMockRepository(t)

	uc := NewUpdateDepartmentUseCase(repo)

	id := int64(1)

	repo.EXPECT().GetByID(id).Return(nil, ErrNotFound)

	result, err := uc.Execute(UpdateDepartmentUseCaseInput{
		ID: id,
	})

	assert.ErrorIs(t, err, ErrNotFound)
	assert.Nil(t, result)
}

func TestUpdateDepartmentUseCase_Execute_ErrInvalidName(t *testing.T) {
	repo := mockdepartment.NewMockRepository(t)

	uc := NewUpdateDepartmentUseCase(repo)

	id := int64(1)
	newName := ""

	dept := &department.Department{ID: id, Name: "Test"}

	repo.EXPECT().GetByID(id).Return(dept, nil)

	result, err := uc.Execute(UpdateDepartmentUseCaseInput{
		ID:   id,
		Name: &newName,
	})

	assert.ErrorIs(t, err, department.ErrInvalidName)
	assert.Nil(t, result)
}

func TestUpdateDepartmentUseCase_Execute_ErrSelfParent(t *testing.T) {
	repo := mockdepartment.NewMockRepository(t)

	uc := NewUpdateDepartmentUseCase(repo)

	id := int64(1)

	dept := &department.Department{ID: id, Name: "Test"}

	repo.EXPECT().GetByID(id).Return(dept, nil)

	result, err := uc.Execute(UpdateDepartmentUseCaseInput{
		ID:       id,
		ParentID: &id,
	})

	assert.ErrorIs(t, err, ErrSelfParent)
	assert.Nil(t, result)
}

func TestUpdateDepartmentUseCase_Execute_ErrCyclicDependency(t *testing.T) {
	repo := mockdepartment.NewMockRepository(t)

	uc := NewUpdateDepartmentUseCase(repo)

	id := int64(1)
	parentID := int64(2)

	dept := &department.Department{ID: id, Name: "Test"}

	repo.EXPECT().GetByID(id).Return(dept, nil)
	repo.EXPECT().GetSubtreeIDs(id).Return([]int64{parentID}, nil)

	result, err := uc.Execute(UpdateDepartmentUseCaseInput{
		ID:       id,
		ParentID: &parentID,
	})

	assert.ErrorIs(t, err, ErrCyclicDependency)
	assert.Nil(t, result)
}

func TestUpdateDepartmentUseCase_Execute_Success(t *testing.T) {
	repo := mockdepartment.NewMockRepository(t)

	uc := NewUpdateDepartmentUseCase(repo)

	id := int64(1)
	parentID := int64(2)

	dept := &department.Department{ID: id, Name: "Test"}
	expectedDept := &department.Department{ID: id, Name: "Test", ParentID: &parentID}

	repo.EXPECT().GetByID(id).Return(dept, nil)
	repo.EXPECT().GetSubtreeIDs(id).Return([]int64{}, nil)
	repo.EXPECT().Update(id, expectedDept).Return(expectedDept, nil)

	result, err := uc.Execute(UpdateDepartmentUseCaseInput{
		ID:       id,
		ParentID: &parentID,
	})

	assert.NoError(t, err)
	assert.Equal(t, &UpdateDepartmentUseCaseOutput{Department: expectedDept}, result)
}
