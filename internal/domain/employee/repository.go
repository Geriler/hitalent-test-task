package employee

type Repository interface {
	Create(employee *Employee) (*Employee, error)
	GetByDepartmentID(departmentID int64) ([]*Employee, error)
	ReassignEmployees(fromDepartmentID, toDepartmentID int64) error
}
