package usecase

import (
	"context"
	"taskMetrics/entity"
	"taskMetrics/internal/user/repository"
	"taskMetrics/pkg/database"
)

type ListUsersUseCase struct {
	process *database.ProcessManager
}

func NewListUsersUseCase(process *database.ProcessManager) *ListUsersUseCase {
	return &ListUsersUseCase{
		process: process,
	}
}

func (u *ListUsersUseCase) Execute(ctx context.Context) ([]entity.User, error) {

	tx, err := u.process.StartProcess(ctx)
	if err != nil {
		return nil, err
	}

	defer u.process.AbortProcess(tx)

	repo := repository.NewUserRepository(tx)

	users, err := repo.List(ctx)
	if err != nil {
		return nil, err
	}

	if err := u.process.CloseProcess(tx); err != nil {
		return nil, err
	}

	return users, nil
}
