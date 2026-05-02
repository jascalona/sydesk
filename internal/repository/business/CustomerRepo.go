package business

import (
	"context"
	"database/sql"
	"sydesk/pkg/domain/business"
)

type customerRepo struct {
	DB *sql.DB
}

func NewCustomerRepo(db *sql.DB) business.CustomerRepo {
	return &customerRepo{DB: db}
}

func (r *customerRepo) GetAll(ctx context.Context) ([]*business.Customer, error) {
	query := `SELECT id, rif, name, created_at FROM customers ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	customers := make([]*business.Customer, 0)
	for rows.Next() {
		customRow := &business.Customer{}
		err := rows.Scan(
			&customRow.ID,
			&customRow.RIF,
			&customRow.NAME,
			&customRow.CREATEDAT,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customRow)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}

// --- REGISTRO DE CLIENTES --- //
func (r *customerRepo) Created(ctx context.Context, customer *business.Customer) error {
	query := `INSERT INTO customers(rif, name) VALUES ($1, $2)`

	_, err := r.DB.ExecContext(ctx, query,
		customer.RIF,
		customer.NAME,
	)
	if err != nil {
		return err
	}
	return nil

}
