package model

import (
	"time"

	"github.com/Geriler/hitalent/internal/domain/department"
)

type Department struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null;size:200"`
	ParentID  *int64 `gorm:"index"`
	CreatedAt time.Time
}

func (d *Department) ToDomain() *department.Department {
	return &department.Department{
		ID:        d.ID,
		Name:      d.Name,
		ParentID:  d.ParentID,
		CreatedAt: d.CreatedAt,
	}
}

func DepartmentFromDomain(d *department.Department) *Department {
	return &Department{
		Name:     d.Name,
		ParentID: d.ParentID,
	}
}
