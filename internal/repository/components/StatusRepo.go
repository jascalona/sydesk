package components

import (
	"context"
	"database/sql"
	"sydesk/pkg/domain/components"
)

type statusRepo struct {
	DB *sql.DB
}

func NewStatusRepo(db *sql.DB) components.StatusRepo {
	return &statusRepo{DB: db}
}

func (r *statusRepo) GetAll(ctx context.Context) ([]*components.Status, error) {

	query := `SELECT id,name,created_at FROM status`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	status := make([]*components.Status, 0)
	for rows.Next() {
		statusRow := &components.Status{}
		err := rows.Scan(
			&statusRow.ID,
			&statusRow.NAME,
			&statusRow.CREATEDAT,
		)

		if err != nil {
			return nil, err
		}
		status = append(status, statusRow)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return status, nil
}
