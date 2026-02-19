package repository

import (
	"context"
	"fmt"
	"taskMetrics/entity"

	"github.com/jmoiron/sqlx"
)

type ListUsersRepository interface {
	List(ctx context.Context) ([]entity.User, error)
}

type listUsersRepo struct {
	db *sqlx.Tx
}

func NewListUsersRepository(db *sqlx.Tx) ListUsersRepository {
	return &listUsersRepo{db: db}
}

func (r *listUsersRepo) List(ctx context.Context) ([]entity.User, error) {

	var users []entity.User
	if err := r.db.SelectContext(ctx, &users, QueryList); err != nil {
		return nil, fmt.Errorf("repository.User.List (query): %w", err)
	}

	return users, nil
}

const QueryList = `SELECT id, username, email, created_at FROM users`
