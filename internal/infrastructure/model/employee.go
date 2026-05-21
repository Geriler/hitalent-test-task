package model

import (
	"time"

	"github.com/Geriler/hitalent/internal/domain/employee"
)

type Employee struct {
	ID           int64  `gorm:"primaryKey;autoIncrement"`
	DepartmentID int64  `gorm:"not null;index"`
	FullName     string `gorm:"not null;size:200"`
	Position     string `gorm:"not null;size:200"`
	HiredAt      *time.Time
	CreatedAt    time.Time
}

func (e *Employee) ToDomain() *employee.Employee {
	return &employee.Employee{
		ID:           e.ID,
		DepartmentID: e.DepartmentID,
		FullName:     e.FullName,
		Position:     e.Position,
		HiredAt:      e.HiredAt,
		CreatedAt:    e.CreatedAt,
	}
}

func EmployeeFromDomain(d *employee.Employee) *Employee {
	return &Employee{
		DepartmentID: d.DepartmentID,
		FullName:     d.FullName,
		Position:     d.Position,
		HiredAt:      d.HiredAt,
	}
}
