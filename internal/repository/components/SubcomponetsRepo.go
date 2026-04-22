package components

import (
	"context"
	"database/sql"
	"sydesk/pkg/domain/components"
)

type subcomponentRepo struct {
	DB *sql.DB
}

func NewSubcomponentRepo(db *sql.DB) components.SubcomponentsRepo {
	return &subcomponentRepo{DB: db}
}

func (r *subcomponentRepo) GetAll(ctx context.Context) ([]*components.Subcomponents, error) {

	query := `SELECT id, name, components_id, description, is_active, created_at FROM subcomponents`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	subcomponents := make([]*components.Subcomponents, 0)
	for rows.Next() {
		subRow := &components.Subcomponents{}
		err := rows.Scan(
			&subRow.ID,
			&subRow.NAME,
			&subRow.COMPONENT_ID,
			&subRow.DESC,
			&subRow.CREATEDAT,
		)
		if err != nil {
			return nil, err
		}
		subcomponents = append(subcomponents, subRow)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return subcomponents, nil

}
