package repository

import (
	"context"
	"taskMetrics/entity"
	"taskMetrics/internal/task"

	"github.com/jmoiron/sqlx"
)

type taskRepo struct {
	create CreateTaskRepository
	update UpdateTaskStatusRepository
	get    GetTaskByIDRepository
}

func NewTaskRepository(tx *sqlx.Tx) task.TaskRepository {
	return &taskRepo{
		create: NewCreateTaskRepository(tx),
		update: NewUpdateTaskStatusRepository(tx),
		get:    NewGetTaskByIDRepository(tx),
	}
}

func (r *taskRepo) Create(ctx context.Context, t *entity.Task) error {
	return r.create.Create(ctx, t)
}

func (r *taskRepo) UpdateStatus(ctx context.Context, id int64, status string) error {
	return r.update.UpdateStatus(ctx, id, status)
}

func (r *taskRepo) GetByID(ctx context.Context, id int64) (*entity.Task, error) {
	return r.get.GetByID(ctx, id)
}
