package business

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain/business"
)

type TicketRequestRepo struct {
	DB *sql.DB
}

func NewTicketRequest(db *sql.DB) business.InterfaceTicketRequest {
	return &TicketRequestRepo{DB: db}
}

func (r *TicketRequestRepo) GetAll(ctx context.Context) ([]*business.TicketRequest, error) {

	query := `
		SELECT 
			id,
			reported_by_customer_id,
			affected_customer_id,
			ticket_bcv,
			enviroment,
			priority,
			type_request,
			sla_id,
			components_id,
			subcomponents_id,
			contact_id,
			topic,
			description,
			created_by,
			created_at,
			expired_in
		FROM ticket_request`

	rows, err := r.DB.QueryContext(ctx, query)

	if err != nil {
		log.Println("ERROR AL CORRER EL QUERY")
		return nil, err
	}

	ticket := make([]*business.TicketRequest, 0)
	for rows.Next() {
		tRow := &business.TicketRequest{}
		err := rows.Scan(
			&tRow.ID,
			&tRow.REPORTED_BY,
			&tRow.AFFECTED_C,
			&tRow.TICKET_BCV,
			&tRow.ENVIROMENT,
			&tRow.PRIORITY,
			&tRow.TYPE_REQUEST,
			&tRow.SLA_ID,
			&tRow.COMPONENT_ID,
			&tRow.SUBCOMPONENT_ID,
			&tRow.CONTACT_ID,
			&tRow.TOPIC,
			&tRow.DESCRIPTION,
			&tRow.CREATED_BY,
			&tRow.CREATED_AT,
			&tRow.EXPIRED_IN,
		)
		if err != nil {
			return nil, err
		}
		ticket = append(ticket, tRow)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ticket, nil

}

func (r *TicketRequestRepo) Created(ctx context.Context, new_ticket *business.TicketRequest) error {
	query := `
		INSERT INTO ticket_request(
			id,
			reported_by_customer_id,
			affected_customer_id,
			ticket_bcv,
			enviroment,
			priority,
			type_request,
			sla_id,
			components_id,
			subcomponents_id,
			contact_id,
			topic,
			description,
			created_by,
			created_at,
			expired_in)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)`

	_, err := r.DB.ExecContext(ctx, query,
		new_ticket.ID,
		new_ticket.REPORTED_BY,
		new_ticket.AFFECTED_C,
		new_ticket.TICKET_BCV,
		new_ticket.ENVIROMENT,
		new_ticket.PRIORITY,
		new_ticket.TYPE_REQUEST,
		new_ticket.SLA_ID,
		new_ticket.COMPONENT_ID,
		new_ticket.SUBCOMPONENT_ID,
		new_ticket.CONTACT_ID,
		new_ticket.TOPIC,
		new_ticket.DESCRIPTION,
		new_ticket.CREATED_BY,
		new_ticket.CREATED_AT,
		new_ticket.EXPIRED_IN,
	)
	if err != nil {
		log.Printf("error al correr el insert")
		return err
	}
	return nil
}
