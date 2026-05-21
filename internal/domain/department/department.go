package department

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrInvalidName = errors.New("invalid department name")
	ErrNotFound    = errors.New("department not found")
)

type Department struct {
	ID        int64
	Name      string
	ParentID  *int64
	CreatedAt time.Time
}

func NewDepartment(name string, parentID *int64) (*Department, error) {
	name, err := validateName(name)
	if err != nil {
		return nil, err
	}

	return &Department{
		Name:     name,
		ParentID: parentID,
	}, nil
}

func (d *Department) Update(name *string, parentID *int64) error {
	if name != nil {
		n, err := validateName(*name)
		if err != nil {
			return err
		}

		d.Name = n
	}
	if parentID != nil {
		d.ParentID = parentID
	}
	return nil
}

func validateName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" || len(name) > 200 {
		return "", ErrInvalidName
	}
	return name, nil
}
