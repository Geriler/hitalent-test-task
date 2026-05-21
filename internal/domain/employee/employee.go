package employee

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrInvalidFullName = errors.New("invalid employee full name")
	ErrInvalidPosition = errors.New("invalid employee position")
)

type Employee struct {
	ID           int64
	DepartmentID int64
	FullName     string
	Position     string
	HiredAt      *time.Time
	CreatedAt    time.Time
}

func NewEmployee(departmentID int64, fullName, position string, hiredAt *time.Time) (*Employee, error) {
	fullName = strings.TrimSpace(fullName)
	if fullName == "" || len(fullName) > 200 {
		return nil, ErrInvalidFullName
	}

	position = strings.TrimSpace(position)
	if position == "" || len(position) > 200 {
		return nil, ErrInvalidPosition
	}

	return &Employee{
		DepartmentID: departmentID,
		FullName:     fullName,
		Position:     position,
		HiredAt:      hiredAt,
	}, nil
}
