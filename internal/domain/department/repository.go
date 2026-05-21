package department

type Repository interface {
	Create(department *Department) (*Department, error)
	GetByID(id int64) (*Department, error)
	Update(id int64, department *Department) (*Department, error)
	Delete(id int64) error

	GetChildren(id int64, depth int) ([]*Department, error)
	ExistsByNameAndParent(name string, parentID *int64) (bool, error)
	GetSubtreeIDs(id int64) ([]int64, error)
}
