package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProcessManager struct {
	db *sqlx.DB
}

func NewProcessManager(db *sqlx.DB) *ProcessManager {
	return &ProcessManager{db: db}
}

func (p *ProcessManager) StartProcess(ctx context.Context) (*sqlx.Tx, error) {
	tenant, err := GetTenant(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	if err := SetSearchPath(ctx, tx, tenant); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("start process: %w", err)
	}

	return tx, nil
}

// CloseProcess = commit
func (p *ProcessManager) CloseProcess(tx *sqlx.Tx) error {
	return tx.Commit()
}

// AbortProcess = rollback
func (p *ProcessManager) AbortProcess(tx *sqlx.Tx) error {
	return tx.Rollback()
}
func SetSearchPath(ctx context.Context, tx *sqlx.Tx, schema string) error {
	if schema == "" {
		return fmt.Errorf("schema cannot be empty")
	}

	_, err := tx.ExecContext(ctx,
		"SELECT set_config('search_path', $1, false)",
		schema,
	)

	if err != nil {
		return fmt.Errorf("SetSearchPath error: %w", err)
	}

	return nil
}
