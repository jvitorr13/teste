package usecase

import (
	"context"
	"fmt"
	"taskMetrics/entity"
	"taskMetrics/internal/task"
	"taskMetrics/internal/task/repository"
	"taskMetrics/pkg/database"

	"github.com/jmoiron/sqlx"
)

type taskUseCase struct {
	db *sqlx.DB
}

func NewTaskUseCase(db *sqlx.DB) task.TaskUseCase {
	return &taskUseCase{db: db}
}

func (u *taskUseCase) CreateTask(ctx context.Context, userID int64, title string) (*entity.Task, error) {
	tenant, err := database.GetTenant(ctx)
	if err != nil {
		return nil, fmt.Errorf("usecase.Task.CreateTask: %w", err)
	}

	tx, err := u.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("usecase.Task.CreateTask (begin tx): %w", err)
	}
	defer tx.Rollback()

	if err := database.SetSearchPath(ctx, tx, tenant); err != nil {
		return nil, fmt.Errorf("usecase.Task.CreateTask (search_path): %w", err)
	}

	repo := repository.NewTaskRepository(tx)
	createUC := NewCreateTaskUseCase(repo)
	res, err := createUC.Execute(ctx, userID, title)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("usecase.Task.CreateTask (commit): %w", err)
	}
	return res, nil
}

func (u *taskUseCase) UpdateTaskStatus(ctx context.Context, id int64, status string) error {
	tenant, err := database.GetTenant(ctx)
	if err != nil {
		return fmt.Errorf("usecase.Task.UpdateTaskStatus: %w", err)
	}

	tx, err := u.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("usecase.Task.UpdateTaskStatus (begin tx): %w", err)
	}
	defer tx.Rollback()

	if err := database.SetSearchPath(ctx, tx, tenant); err != nil {
		return fmt.Errorf("usecase.Task.UpdateTaskStatus (search_path): %w", err)
	}

	repo := repository.NewTaskRepository(tx)
	updateUC := NewUpdateTaskStatusUseCase(repo)
	if err := updateUC.Execute(ctx, id, status); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("usecase.Task.UpdateTaskStatus (commit): %w", err)
	}
	return nil
}
