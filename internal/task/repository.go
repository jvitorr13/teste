package task

import (
	"context"
	"taskMetrics/entity"
)

type TaskRepository interface {
	Create(ctx context.Context, task *entity.Task) error
	UpdateStatus(ctx context.Context, id int64, status string) error
	GetByID(ctx context.Context, id int64) (*entity.Task, error)
}
