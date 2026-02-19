package usecase

import (
	"context"
	"taskMetrics/entity"
	"taskMetrics/internal/task"
)

type CreateTaskUseCase interface {
	Execute(ctx context.Context, userID int64, title string) (*entity.Task, error)
}

type createTaskUseCase struct {
	taskRepo task.TaskRepository
}

func NewCreateTaskUseCase(repo task.TaskRepository) CreateTaskUseCase {
	return &createTaskUseCase{taskRepo: repo}
}

func (u *createTaskUseCase) Execute(ctx context.Context, userID int64, title string) (*entity.Task, error) {
	task := &entity.Task{UserID: userID, Title: title, Status: "PENDING"}
	if err := u.taskRepo.Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}
