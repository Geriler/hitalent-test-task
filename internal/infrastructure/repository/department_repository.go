package repository

import (
	"errors"

	"github.com/Geriler/hitalent/internal/domain/department"
	"github.com/Geriler/hitalent/internal/infrastructure/model"
	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{
		db: db,
	}
}

func (r *DepartmentRepository) Create(d *department.Department) (*department.Department, error) {
	m := model.DepartmentFromDomain(d)

	result := r.db.Create(&m)
	if result.Error != nil {
		return nil, result.Error
	}

	return m.ToDomain(), nil
}

func (r *DepartmentRepository) GetByID(id int64) (*department.Department, error) {
	m := &model.Department{}

	result := r.db.First(m, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, department.ErrNotFound
		}
		return nil, result.Error
	}

	return m.ToDomain(), nil
}

func (r *DepartmentRepository) Update(id int64, d *department.Department) (*department.Department, error) {
	m := model.DepartmentFromDomain(d)
	m.ID = id

	result := r.db.Updates(&m)
	if result.Error != nil {
		return nil, result.Error
	}

	updated := &model.Department{}
	result = r.db.First(updated, m.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return updated.ToDomain(), nil
}

func (r *DepartmentRepository) Delete(id int64) error {
	m := &model.Department{}

	result := r.db.Delete(m, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *DepartmentRepository) GetChildren(id int64, depth int) ([]*department.Department, error) {
	var models []*model.Department

	result := r.db.Raw(`
		WITH RECURSIVE subtree AS (
			SELECT *, 1 AS depth FROM departments WHERE parent_id = ?
			UNION ALL
			SELECT d.*, s.depth + 1 FROM departments d
			JOIN subtree s ON d.parent_id = s.id
			WHERE s.depth < ?
		)
		SELECT * FROM subtree
	`, id, depth).Scan(&models)
	if result.Error != nil {
		return nil, result.Error
	}

	departments := make([]*department.Department, len(models))
	for i, d := range models {
		departments[i] = d.ToDomain()
	}

	return departments, nil
}

func (r *DepartmentRepository) ExistsByNameAndParent(name string, parentID *int64) (bool, error) {
	m := &model.Department{}

	var result *gorm.DB
	if parentID != nil {
		result = r.db.Where("name = ? and parent_id = ?", name, *parentID).First(m)
	} else {
		result = r.db.Where("name = ? and parent_id is null", name).First(m)
	}

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

func (r *DepartmentRepository) GetSubtreeIDs(id int64) ([]int64, error) {
	var ids []int64

	result := r.db.Raw(`
		WITH RECURSIVE subtree AS (
			SELECT id FROM departments WHERE parent_id = ?
			UNION ALL
			SELECT d.id FROM departments d
			JOIN subtree s ON d.parent_id = s.id
		)
		SELECT id FROM subtree
	`, id).Scan(&ids)

	if result.Error != nil {
		return nil, result.Error
	}

	return ids, nil
}
