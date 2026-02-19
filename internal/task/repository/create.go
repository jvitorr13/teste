package repository

import (
	"context"
	"fmt"
	"taskMetrics/entity"

	"github.com/jmoiron/sqlx"
)

type CreateTaskRepository interface {
	Create(ctx context.Context, t *entity.Task) error
}

type createTaskRepo struct {
	db *sqlx.Tx
}

func NewCreateTaskRepository(db *sqlx.Tx) CreateTaskRepository {
	return &createTaskRepo{db: db}
}

func (r *createTaskRepo) Create(ctx context.Context, t *entity.Task) error {
	query := `
 		INSERT INTO tasks (title, status, user_id) 
 		VALUES ($1, $2, $3) 
 		RETURNING id, created_at, updated_at
 	`

	err := r.db.QueryRowxContext(ctx, query, t.Title, t.Status, t.UserID).
		Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return fmt.Errorf("repository.Task.Create (query): %w", err)
	}

	return nil
}
