package business

import (
	"context"
	"database/sql"
	"log"
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

// --- MODULO AUDITORIA ---//
func (r *customerRepo) GetAuditCustomer(ctx context.Context, customer_id int) ([]*business.Customer, error) {
	query := `
		SELECT
			cli.id,
			cli.rif,
			cli.name,
			cli.channel,
			cli.ws_group,
			cli.uid_sypago,
			cli.created_at,
			cpr.product_id,
			cpr.id AS product_role_id,
			ibp.name AS nombre_banco,
			cont.name AS contact_name,
			cont.surname AS contact_surname,
			cont.email AS contact_email,
			cont.phone AS contact_phone
		FROM customer_product_roles cpr
		INNER JOIN customers cli ON cpr.customer_id = cli.id
		INNER JOIN contact cont ON cont.customer_id = cli.id
		INNER JOIN customer_product_roles prod_padre ON cpr.parent_id = prod_padre.id
		INNER JOIN customers ibp ON prod_padre.customer_id = ibp.id
		WHERE cli.id != ibp.id
			AND ($1 = 0 OR cli.id = $1)
		ORDER BY cli.created_at DESC
	`

	rows, err := r.DB.QueryContext(ctx, query, customer_id)

	if err != nil {
		log.Printf("Error al correr el query context: %v", err)
		return nil, err
	}
	defer rows.Close()

	customers := make([]*business.Customer, 0)
	for rows.Next() {
		customRow := &business.Customer{}
		err := rows.Scan(
			&customRow.ID,
			&customRow.RIF,
			&customRow.NAME,
			&customRow.CHANNEL,
			&customRow.WS_GROUP,
			&customRow.UID_SYPAGO,
			&customRow.CREATEDAT,
			&customRow.PRODUCT_ID,
			&customRow.PRODUCT_ROLE,
			&customRow.NOMBRE_BANCO,
			&customRow.CONTACT_NAME,
			&customRow.CONTACT_LAST,
			&customRow.CONTACT_EMAIL,
			&customRow.CONTACT_PHONE,
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
