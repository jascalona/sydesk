package components

import (
	"context"
	"database/sql"
	"sydesk/pkg/domain/components"
)

type priorityRepo struct {
	DB *sql.DB
}

func NewPriorityRepo(db *sql.DB) components.InterfacePriority {
	return &priorityRepo{DB: db}
}

func (r *priorityRepo) GetAll(ctx context.Context) ([]*components.Priority, error) {
	query := `SELECT id, name, created_at FROM priority`

	rows, err := r.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	priority := make([]*components.Priority, 0)

	for rows.Next() {
		prRow := &components.Priority{}
		err := rows.Scan(
			&prRow.ID,
			&prRow.NAME,
			&prRow.CREATEDAT,
		)
		if err != nil {
			return nil, err
		}
		priority = append(priority, prRow)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return priority, nil

}
