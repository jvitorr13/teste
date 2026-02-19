package repository

import (
	"context"
	"fmt"
	"taskMetrics/entity"

	"github.com/jmoiron/sqlx"
)

type CreateUserRepo struct {
	db *sqlx.Tx
}

func NewCreateUserRepository(db *sqlx.Tx) CreateUserRepo {
	return &CreateUserRepo{db: db}
}

func (r *CreateUserRepo) Create(ctx context.Context, u *entity.User) error {

	err := r.db.QueryRowxContext(ctx, QueryCreate, u.Username, u.Email).Scan(&u.ID, &u.CreatedAt)
	if err != nil {
		return fmt.Errorf("repository.User.Create (query): %w", err)
	}

	return nil
}

const QueryCreate = `
 		INSERT INTO users (username, email) 
 		VALUES ($1, $2) 
 		RETURNING id, created_at
 	`
