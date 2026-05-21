package repository

import (
	"github.com/Geriler/hitalent/internal/domain/employee"
	"github.com/Geriler/hitalent/internal/infrastructure/model"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (r *EmployeeRepository) Create(e *employee.Employee) (*employee.Employee, error) {
	m := model.EmployeeFromDomain(e)

	result := r.db.Create(&m)
	if result.Error != nil {
		return nil, result.Error
	}

	return m.ToDomain(), nil
}

func (r *EmployeeRepository) GetByDepartmentID(departmentID int64) ([]*employee.Employee, error) {
	var models []*model.Employee

	result := r.db.Where("department_id = ?", departmentID).Find(&models)
	if result.Error != nil {
		return nil, result.Error
	}

	employees := make([]*employee.Employee, len(models))
	for i, m := range models {
		employees[i] = m.ToDomain()
	}

	return employees, nil
}

func (r *EmployeeRepository) ReassignEmployees(fromDepartmentID, toDepartmentID int64) error {
	result := r.db.Model(&model.Employee{}).
		Where("department_id = ?", fromDepartmentID).
		Update("department_id", toDepartmentID)
	return result.Error
}
