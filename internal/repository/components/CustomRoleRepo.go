package components

import (
	"context"
	"database/sql"
	"sydesk/pkg/domain/components"
)

type customRoleRepo struct {
	DB *sql.DB
}

func NewCustomRoleRepo(db *sql.DB) components.CustomRoleRepo {
	return &customRoleRepo{DB: db}
}

func (r *customRoleRepo) GetAll(ctx context.Context) ([]*components.CustomerRole, error) {
	query := `SELECT id,name,created_at FROM customer_role`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	customRole := make([]*components.CustomerRole, 0)

	for rows.Next() {
		crRow := &components.CustomerRole{}
		err := rows.Scan(
			&crRow.ID,
			&crRow.NAME,
			&crRow.CREATEDAT,
		)
		if err != nil {
			return nil, err
		}
		customRole = append(customRole, crRow)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return customRole, nil
}
