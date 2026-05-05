package components

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain/components"
)

type customerPRrepo struct {
	DB *sql.DB
}

func NewCustomProductRoleRepo(db *sql.DB) components.CustomerProductRoleRepo {
	return &customerPRrepo{DB: db}
}

func (r *customerPRrepo) GetAll(ctx context.Context) ([]*components.CustomerProductRole, error) {

	query := `SELECT id, customer_id, product_id, role_id, parent_id, created_at FROM customer_product_roles`

	rows, err := r.DB.QueryContext(ctx, query)

	if err != nil {
		log.Printf("error al correr el query context", err.Error())
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
			&cpr.PARENT_ID,
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

func (r *customerPRrepo) Created(ctx context.Context, cpr *components.CustomerProductRole) error {

	query := `
		INSERT INTO customer_product_roles (
			customer_id,
			product_id,
			role_id,
			parent_id)VALUES($1,$2,$3,$4)`

	_, err := r.DB.ExecContext(ctx, query,
		cpr.CUSTOMER_ID,
		cpr.PRODUCT_ID,
		cpr.ROLE_ID,
		cpr.PARENT_ID,
	)
	if err != nil {
		log.Printf("error al corrrer el insert")
		return err
	}
	return nil
}
