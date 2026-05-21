package handler

import (
	"time"

	"github.com/Geriler/hitalent/internal/domain/department"
	"github.com/Geriler/hitalent/internal/domain/employee"
)

type DepartmentDTO struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	ParentID  *int64 `json:"parent_id"`
	CreatedAt string `json:"created_at"`
}

type EmployeeDTO struct {
	ID           int64   `json:"id"`
	DepartmentID int64   `json:"department_id"`
	FullName     string  `json:"full_name"`
	Position     string  `json:"position"`
	HiredAt      *string `json:"hired_at"`
	CreatedAt    string  `json:"created_at"`
}

func toDepartmentDTO(d *department.Department) *DepartmentDTO {
	return &DepartmentDTO{
		ID:        d.ID,
		Name:      d.Name,
		ParentID:  d.ParentID,
		CreatedAt: d.CreatedAt.Format(time.RFC3339),
	}
}

func toDepartmentDTOs(departments []*department.Department) []*DepartmentDTO {
	dtos := make([]*DepartmentDTO, len(departments))
	for i, d := range departments {
		dtos[i] = toDepartmentDTO(d)
	}
	return dtos
}

func toEmployeeDTO(e *employee.Employee) *EmployeeDTO {
	var hiredAt *string

	if e.HiredAt != nil {
		t := e.HiredAt.Format(time.RFC3339)
		hiredAt = &t
	}

	return &EmployeeDTO{
		ID:           e.ID,
		DepartmentID: e.DepartmentID,
		FullName:     e.FullName,
		Position:     e.Position,
		HiredAt:      hiredAt,
		CreatedAt:    e.CreatedAt.Format(time.RFC3339),
	}
}

func toEmployeeDTOs(employees []*employee.Employee) []*EmployeeDTO {
	dtos := make([]*EmployeeDTO, len(employees))
	for i, e := range employees {
		dtos[i] = toEmployeeDTO(e)
	}
	return dtos
}
