package user

import (
	"context"
	"taskMetrics/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	List(ctx context.Context) ([]entity.User, error)
}
