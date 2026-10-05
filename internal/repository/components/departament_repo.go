package components

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain/components"
)

type departamentRepo struct {
	DB *sql.DB
}

func NewDepartamentRepo(db *sql.DB) *departamentRepo {
	return &departamentRepo{DB: db}
}

func (r *departamentRepo) GetAll(ctx context.Context) ([]*components.Departament, error) {
	query := "SELECT id, name, is_active, created_at FROM departament"
	rows, err := r.DB.QueryContext(ctx, query)

	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		return nil, err
	}

	defer rows.Close()

	departaments := make([]*components.Departament, 0)

	for rows.Next() {
		departament := &components.Departament{}
		err := rows.Scan(
			&departament.ID,
			&departament.NAME,
			&departament.IS_ACTIVE,
			&departament.CREATED_AT,
		)
		if err != nil {
			log.Printf("Error al escanear la fila: %v", err)
			return nil, err
		}
		departaments = append(departaments, departament)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departaments, nil
}
