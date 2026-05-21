package handler_test

import (
	"context"
	"testing"
	"time"

	"github.com/Geriler/hitalent/internal/infrastructure/database"
	infrahttp "github.com/Geriler/hitalent/internal/infrastructure/http"
	"github.com/Geriler/hitalent/internal/infrastructure/http/handler"
	"github.com/Geriler/hitalent/internal/infrastructure/repository"
	"github.com/Geriler/hitalent/internal/usecase/department"
	"github.com/Geriler/hitalent/internal/usecase/employee"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
)

func setupTestContainer(t *testing.T) *gorm.DB {
	t.Helper()

	ctx := context.Background()

	container, err := postgres.Run(ctx, "postgres:18-alpine",
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Minute),
		),
	)
	if err != nil {
		t.Fatalf("failed to start postgres container: %v", err)
	}

	connStr := container.MustConnectionString(ctx, "sslmode=disable")
	dbGORM, err := database.NewPostgresDB(connStr, true)
	if err != nil {
		t.Fatalf("failed to connect to postgres: %v", err)
	}

	db, err := dbGORM.DB()
	if err != nil {
		t.Fatalf("failed to connect to postgres: %v", err)
	}

	err = goose.Up(db, "../../../../migrations")
	if err != nil {
		t.Fatalf("failed to run migrations: %v", err)
	}

	t.Cleanup(func() {
		err = db.Close()
		if err != nil {
			t.Fatalf("failed to close postgres connection: %v", err)
		}
		err = container.Terminate(ctx)
		if err != nil {
			t.Fatalf("failed to terminate postgres container: %v", err)
		}
	})

	return dbGORM
}

func setupRouter(t *testing.T) *infrahttp.Router {
	t.Helper()

	db := setupTestContainer(t)

	departmentRepo := repository.NewDepartmentRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)

	departmentCreateUC := department.NewCreateDepartmentUseCase(departmentRepo)
	departmentGetUC := department.NewGetDepartmentUseCase(departmentRepo, employeeRepo)
	departmentUpdateUC := department.NewUpdateDepartmentUseCase(departmentRepo)
	departmentDeleteUC := department.NewDeleteDepartmentUseCase(departmentRepo, employeeRepo)

	employeeCreateUC := employee.NewCreateEmployeeUseCase(employeeRepo, departmentRepo)

	departmentHandler := handler.NewDepartmentHandler(departmentCreateUC, departmentGetUC, departmentUpdateUC, departmentDeleteUC)
	employeeHandler := handler.NewEmployeeHandler(employeeCreateUC)

	return infrahttp.NewRouter(departmentHandler, employeeHandler)
}
