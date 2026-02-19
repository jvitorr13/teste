package user

import (
	"context"
	"taskMetrics/entity"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	ListUsers(ctx context.Context) ([]entity.User, error)
}
