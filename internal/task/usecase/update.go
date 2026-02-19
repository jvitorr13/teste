package usecase

import (
	"context"
	"taskMetrics/internal/task"
)

type UpdateTaskStatusUseCase interface {
	Execute(ctx context.Context, id int64, status string) error
}

type updateTaskStatusUseCase struct {
	taskRepo task.TaskRepository
}

func NewUpdateTaskStatusUseCase(repo task.TaskRepository) UpdateTaskStatusUseCase {
	return &updateTaskStatusUseCase{taskRepo: repo}
}

func (u *updateTaskStatusUseCase) Execute(ctx context.Context, id int64, status string) error {
	return u.taskRepo.UpdateStatus(ctx, id, status)
}
