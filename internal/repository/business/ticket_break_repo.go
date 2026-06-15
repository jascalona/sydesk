package business

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain/business"
)

type TicketBreakRepo struct {
	DB *sql.DB
}

func NewTicketBreakRepo(db *sql.DB) business.InterfaceBreak {
	return &TicketBreakRepo{DB: db}
}

func (r *TicketBreakRepo) GetAll(ctx context.Context) ([]*business.TicketBreak, error) {
	query := `
		SELECT 
			id,
			ticket_id,
			status_paused,
			start_at,
			end_at, 
			created_at
		FROM ticket_break ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query)

	if err != nil {
		log.Println("Error al correr el query")
		return nil, err
	}

	ticket_break := make([]*business.TicketBreak, 0)

	for rows.Next() {
		tRB := &business.TicketBreak{}
		err := rows.Scan(
			&tRB.ID,
			&tRB.TICKET_ID,
			&tRB.STS_P,
			&tRB.START_AT,
			&tRB.END_AT,
			&tRB.CREATED_AT,
		)
		if err != nil {
			return nil, err
		}

		ticket_break = append(ticket_break, tRB)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ticket_break, nil
}
