package components

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain/components"
)

type componentsRepo struct {
	DB *sql.DB
}

func NewComponentsRepo(db *sql.DB) components.ComponentsRepo {
	return &componentsRepo{DB: db}
}

func (r *componentsRepo) GetAll(ctx context.Context) ([]*components.Components, error) {

	query := `SELECT id, name, product_id, description, is_active, created_at FROM components`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	component := make([]*components.Components, 0)
	for rows.Next() {
		componentRow := &components.Components{}
		err := rows.Scan(
			&componentRow.ID,
			&componentRow.NAME,
			&componentRow.PRODUCT_ID,
			&componentRow.DESC,
			&componentRow.IS_ACTIVE,
			&componentRow.CREATEDAT,
		)
		if err != nil {
			return nil, err
		}
		component = append(component, componentRow)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return component, nil
}

func (r *componentsRepo) Created(ctx context.Context, component *components.Components) error {
	query := "INSERT INTO components (name, product_id, description, is_active)VALUES($1, $2, $3, $4)"

	_, err := r.DB.ExecContext(ctx, query,
		component.NAME,
		component.PRODUCT_ID,
		component.DESC,
		component.IS_ACTIVE,
	)
	if err != nil {
		log.Printf("Error al invocar el ExceContext", err.Error())
		return err
	}
	return nil

}
