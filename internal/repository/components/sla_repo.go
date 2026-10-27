package components

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain/components"
)

type SLARepo struct {
	DB *sql.DB
}

func NewSLARepo(db *sql.DB) components.InterfaceSLA {
	return &SLARepo{DB: db}
}

func (r *SLARepo) GetAll(ctx context.Context) ([]*components.SLA, error) {

	query := `SELECT id, name, unit_time_id, amount, created_at FROM sla`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	slas := make([]*components.SLA, 0)
	for rows.Next() {
		slaRow := &components.SLA{}
		err := rows.Scan(
			&slaRow.ID,
			&slaRow.NAME,
			&slaRow.UNIT_TIME_ID,
			&slaRow.AMOUNT,
			&slaRow.CREATED_AT,
		)
		if err != nil {
			return nil, err
		}
		slas = append(slas, slaRow)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al escanear los registros: ", err)
	}

	return slas, nil

}
