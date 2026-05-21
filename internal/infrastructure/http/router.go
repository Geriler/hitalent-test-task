package http

import (
	"net/http"

	"github.com/Geriler/hitalent/internal/infrastructure/http/handler"
)

type Router struct {
	departmentHandler *handler.DepartmentHandler
	employeeHandler   *handler.EmployeeHandler
}

func NewRouter(departmentHandler *handler.DepartmentHandler, employeeHandler *handler.EmployeeHandler) *Router {
	return &Router{
		departmentHandler: departmentHandler,
		employeeHandler:   employeeHandler,
	}
}

func (r *Router) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /departments", r.departmentHandler.Create)
	mux.HandleFunc("GET /departments/{id}", r.departmentHandler.Get)
	mux.HandleFunc("PATCH /departments/{id}", r.departmentHandler.Update)
	mux.HandleFunc("DELETE /departments/{id}", r.departmentHandler.Delete)
	mux.HandleFunc("POST /departments/{id}/employees", r.employeeHandler.Create)

	return mux
}
