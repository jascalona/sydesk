package components

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain/components"
)

type TypeRequestRepo struct {
	DB *sql.DB
}

func NewTypeRequestRepo(db *sql.DB) components.InterfaceTypeRequest {
	return &TypeRequestRepo{DB: db}
}

func (r *TypeRequestRepo) GetAll(ctx context.Context) ([]*components.TypeRequest, error) {
	query := `SELECT id, name, created_at FROM type_request ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("ERROR AL CORRER EL QUERY")
		return nil, err
	}

	tyrequest := make([]*components.TypeRequest, 0)
	for rows.Next() {
		tpR := &components.TypeRequest{}
		err := rows.Scan(
			&tpR.ID,
			&tpR.NAME,
			&tpR.CREATEDAT,
		)

		if err != nil {
			return nil, err
		}

		tyrequest = append(tyrequest, tpR)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tyrequest, nil

}
