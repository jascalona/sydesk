package business

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain/business"
)

type contactRepo struct {
	DB *sql.DB
}

func NewContactRepo(db *sql.DB) business.InterfaceContact {
	return &contactRepo{DB: db}
}

func (r *contactRepo) GetAll(ctx context.Context) ([]*business.Contact, error) {
	query :=
		`SELECT 
				id,
				customer_id,
				name,
				surname,
				email,
				phone,
				created_at
			FROM contact ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query)

	if err != nil {
		log.Println("error al correr el query", err.Error())
		return nil, err
	}

	contact := make([]*business.Contact, 0)

	for rows.Next() {
		rc := &business.Contact{}
		err := rows.Scan(
			rc.ID,
			rc.CUSTOMER_ID,
			rc.NAME,
			rc.SURNAME,
			rc.EMAIL,
			rc.PHONE,
			rc.CREATED_AT,
		)
		if err != nil {
			return nil, err
		}

		contact = append(contact, rc)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return contact, nil

}

func (r *contactRepo) Created(ctx context.Context, contact *business.Contact) error {
	query := `insert into contact (customer_id,name,surname,email,phone)VALUES($1,$2,$3,$4,$5)`

	_, err := r.DB.ExecContext(ctx, query,
		contact.CUSTOMER_ID,
		contact.NAME,
		contact.SURNAME,
		contact.EMAIL,
		contact.PHONE,
	)

	if err != nil {
		log.Println("error al correr el query", err.Error())
		return err
	}
	return nil
}
