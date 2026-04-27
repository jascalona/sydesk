package components

import (
	"context"
	"database/sql"
	"sydesk/pkg/domain/components"
)

type customerPRrepo struct {
	DB *sql.DB
}

func NewCustomProductRoleRepo(db *sql.DB) components.CustomerProductRoleRepo {
	return &customerPRrepo{DB: db}
}

func (r *customerPRrepo) GetAll(ctx context.Context) ([]*components.CustomerProductRole, error) {

	query := `SELECT id, customer_id, product_id, role_id, created_at FROM customer_product_roles`

	rows, err := r.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	customProductRole := make([]*components.CustomerProductRole, 0)

	for rows.Next() {
		cpr := &components.CustomerProductRole{}
		err := rows.Scan(
			&cpr.ID,
			&cpr.CUSTOMER_ID,
			&cpr.PRODUCT_ID,
			&cpr.ROLE_ID,
			&cpr.CREATEDAT,
		)
		if err != nil {
			return nil, err
		}
		customProductRole = append(customProductRole, cpr)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return customProductRole, nil

}
