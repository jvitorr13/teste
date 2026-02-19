package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type UpdateTaskStatusRepository interface {
	UpdateStatus(ctx context.Context, id int64, status string) error
}

type updateTaskStatusRepo struct {
	db *sqlx.Tx
}

func NewUpdateTaskStatusRepository(db *sqlx.Tx) UpdateTaskStatusRepository {
	return &updateTaskStatusRepo{db: db}
}

func (r *updateTaskStatusRepo) UpdateStatus(ctx context.Context, id int64, status string) error {
	query := `UPDATE tasks SET status = $1, updated_at = $2 WHERE id = $3`

	_, err := r.db.ExecContext(ctx, query, status, time.Now(), id)
	if err != nil {
		return fmt.Errorf("repository.Task.UpdateStatus (query): %w", err)
	}

	return nil
}
