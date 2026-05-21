package department

import (
	"testing"

	"github.com/Geriler/hitalent/internal/domain/department"
	mockdepartment "github.com/Geriler/hitalent/internal/mocks/department"
	"github.com/stretchr/testify/assert"
)

func TestCreateDepartmentUseCase_Execute_ErrInvalidName(t *testing.T) {
	repo := mockdepartment.NewMockRepository(t)

	uc := NewCreateDepartmentUseCase(repo)

	result, err := uc.Execute(CreateDepartmentUseCaseInput{
		Name: "",
	})

	assert.ErrorIs(t, err, department.ErrInvalidName)
	assert.Nil(t, result)
}

func TestCreateDepartmentUseCase_Execute_ErrParentNotFound(t *testing.T) {
	repo := mockdepartment.NewMockRepository(t)

	name := "Test Department"
	parentID := int64(1)

	repo.EXPECT().GetByID(parentID).Return(nil, ErrParentNotFound)

	uc := NewCreateDepartmentUseCase(repo)

	result, err := uc.Execute(CreateDepartmentUseCaseInput{
		Name:     name,
		ParentID: &parentID,
	})

	assert.ErrorIs(t, err, ErrParentNotFound)
	assert.Nil(t, result)
}

func TestCreateDepartmentUseCase_Execute_ErrDepartmentAlreadyExist(t *testing.T) {
	repo := mockdepartment.NewMockRepository(t)

	name := "Test Department"
	parentID := int64(1)
	parentName := "Test Parent"

	repo.EXPECT().GetByID(parentID).Return(&department.Department{ID: parentID, Name: parentName}, nil)
	repo.EXPECT().ExistsByNameAndParent(name, &parentID).Return(true, nil)

	uc := NewCreateDepartmentUseCase(repo)

	result, err := uc.Execute(CreateDepartmentUseCaseInput{
		Name:     name,
		ParentID: &parentID,
	})

	assert.ErrorIs(t, err, ErrAlreadyExists)
	assert.Nil(t, result)
}

func TestCreateDepartmentUseCase_Execute_Success(t *testing.T) {
	repo := mockdepartment.NewMockRepository(t)

	name := "Test Department"

	expected := &department.Department{ID: 1, Name: name}

	repo.EXPECT().ExistsByNameAndParent(name, (*int64)(nil)).Return(false, nil)
	repo.EXPECT().Create(&department.Department{Name: name, ParentID: nil}).Return(expected, nil)

	uc := NewCreateDepartmentUseCase(repo)

	result, err := uc.Execute(CreateDepartmentUseCaseInput{
		Name: name,
	})

	assert.NoError(t, err)
	assert.Equal(t, expected, result.Department)
}
