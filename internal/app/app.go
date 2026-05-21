package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/Geriler/hitalent/internal/config"
	"github.com/Geriler/hitalent/internal/infrastructure/database"
	infrahttp "github.com/Geriler/hitalent/internal/infrastructure/http"
	"github.com/Geriler/hitalent/internal/infrastructure/http/handler"
	"github.com/Geriler/hitalent/internal/infrastructure/repository"
	"github.com/Geriler/hitalent/internal/usecase/department"
	"github.com/Geriler/hitalent/internal/usecase/employee"
)

type App struct {
	logger *slog.Logger
	server *http.Server
	router *infrahttp.Router
}

func NewApp(logger *slog.Logger) *App {
	a := &App{
		logger: logger,
	}

	a.initDeps()

	return a
}

func (a *App) initDeps() {
	deps := []func(){
		a.initRouter,
		a.initHTTPServer,
	}
	for _, fn := range deps {
		fn()
	}
}

func (a *App) initRouter() {
	cfg := config.MustGet()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name,
	)

	db, err := database.NewPostgresDB(dsn, false)
	if err != nil {
		a.logger.Error("db init error", "err", err)
		os.Exit(1)
	}

	departmentRepo := repository.NewDepartmentRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)

	departmentCreateUC := department.NewCreateDepartmentUseCase(departmentRepo)
	departmentGetUC := department.NewGetDepartmentUseCase(departmentRepo, employeeRepo)
	departmentUpdateUC := department.NewUpdateDepartmentUseCase(departmentRepo)
	departmentDeleteUC := department.NewDeleteDepartmentUseCase(departmentRepo, employeeRepo)
	departmentHandler := handler.NewDepartmentHandler(departmentCreateUC, departmentGetUC, departmentUpdateUC, departmentDeleteUC)

	employeeCreateUC := employee.NewCreateEmployeeUseCase(employeeRepo, departmentRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeCreateUC)

	router := infrahttp.NewRouter(departmentHandler, employeeHandler)

	a.router = router
}

func (a *App) initHTTPServer() {
	cfg := config.MustGet()

	a.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
		Handler: a.router.Routes(),
	}
}

func (a *App) Start(ctx context.Context) error {
	a.logger.Info("starting server")
	return a.server.ListenAndServe()
}

func (a *App) Stop(ctx context.Context) error {
	a.logger.Info("stopping server")
	return a.server.Shutdown(ctx)
}
