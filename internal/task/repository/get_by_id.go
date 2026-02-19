package repository

import (
	"context"
	"fmt"
	"taskMetrics/entity"

	"github.com/jmoiron/sqlx"
)

type GetTaskByIDRepository interface {
	GetByID(ctx context.Context, id int64) (*entity.Task, error)
}

type getTaskByIDRepo struct {
	db *sqlx.Tx
}

func NewGetTaskByIDRepository(db *sqlx.Tx) GetTaskByIDRepository {
	return &getTaskByIDRepo{db: db}
}

func (r *getTaskByIDRepo) GetByID(ctx context.Context, id int64) (*entity.Task, error) {
	query := `
 		SELECT id, title, status, user_id, created_at, updated_at 
 		FROM tasks 
 		WHERE id = $1
 	`

	var t entity.Task
	if err := r.db.GetContext(ctx, &t, query, id); err != nil {
		return nil, fmt.Errorf("repository.Task.GetByID (query): %w", err)
	}

	return &t, nil
}
