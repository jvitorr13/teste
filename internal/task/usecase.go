package task

import (
	"context"
	"taskMetrics/entity"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, userID int64, title string) (*entity.Task, error)
	UpdateTaskStatus(ctx context.Context, id int64, status string) error
}
