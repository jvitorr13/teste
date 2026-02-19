package usecase

import (
	"context"
	"taskMetrics/entity"
	"taskMetrics/internal/user/repository"
	"taskMetrics/pkg/database"
)

type CreateUserUseCase struct {
	process *database.ProcessManager
}

func NewCreateUserUseCase(process *database.ProcessManager) *CreateUserUseCase {
	return &CreateUserUseCase{
		process: process,
	}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, user *entity.User) (*entity.User, error) {
	tx, err := u.process.StartProcess(ctx)
	if err != nil {
		return nil, err
	}

	defer u.process.AbortProcess(tx)

	repo := repository.NewUserRepository(tx)

	if user, err = repo.Create(ctx, user); err != nil {
		return nil, err
	}

	if err := u.process.CloseProcess(tx); err != nil {
		return nil, err
	}

	return user, nil
}
