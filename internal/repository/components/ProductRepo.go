package components

import (
	"context"
	"database/sql"
	"sydesk/pkg/domain/components"
)

type productRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) components.ProductRepo {
	return &productRepo{DB: db}
}

func (r *productRepo) GetAll(ctx context.Context) ([]*components.Product, error) {
	query := `SELECT id, name, description, is_active, created_at FROM product`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	products := make([]*components.Product, 0)
	for rows.Next() {
		productRow := &components.Product{}
		err := rows.Scan(
			&productRow.ID,
			&productRow.NAME,
			&productRow.DESC,
			&productRow.IS_ACTIVE,
			&productRow.CREATEDAT,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, productRow)
	}
	// validacion error
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
