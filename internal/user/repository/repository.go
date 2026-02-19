package repository

import (
	"context"
	"taskMetrics/entity"
	"taskMetrics/internal/user"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	create CreateUserRepo
	list   ListUsersRepository
}

func NewUserRepository(tx *sqlx.Tx) user.UserRepository {
	return &userRepo{
		create: NewCreateUserRepository(tx),
		list:   NewListUsersRepository(tx),
	}
}

func (r *userRepo) Create(ctx context.Context, u *entity.User) error {
	return r.create.Create(ctx, u)
}

func (r *userRepo) List(ctx context.Context) ([]entity.User, error) {
	return r.list.List(ctx)
}
