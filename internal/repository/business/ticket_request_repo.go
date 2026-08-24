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

// VW CONSUMER REQUEST
func (r *TicketRequestRepo) GetVWTicketRequest(ctx context.Context) ([]*business.VWTicketRequest, error) {
	query := `
		SELECT 
			id,
			ticket_bcv,
			enviroment,
			type_request,
			topic,
			description,
			created_by,
			created_at,
			expired_in,
			priority,
			affected_customer,
			sla,
			component,
			subcomponent,
			product,
			contact,
			current_status
		FROM vw_ticket_request_details ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query)

	if err != nil {
		log.Println("ERROR AL CORRER EL QUERY: ", err.Error())
		return nil, err
	}

	vw_request := make([]*business.VWTicketRequest, 0)

	for rows.Next() {
		vw_r := &business.VWTicketRequest{}
		err := rows.Scan(
			&vw_r.ID,
			&vw_r.TICKET_BCV,
			&vw_r.ENVIROMENT,
			&vw_r.TYPE_REQUEST,
			&vw_r.TOPIC,
			&vw_r.DESCRIPTION,
			&vw_r.CREATED_BY,
			&vw_r.CREATED_AT,
			&vw_r.EXPIRED_IN,
			&vw_r.PRIORITY,
			&vw_r.AFFECTED_C,
			&vw_r.SLA,
			&vw_r.COMPONENT,
			&vw_r.SUBCOMPONENT,
			&vw_r.PRODUCT,
			&vw_r.CONTACT,
			&vw_r.CURRENT_STATUS,
		)

		if err != nil {
			log.Println("Error al aplicar el escaner: ", err.Error())
			return nil, err
		}

		vw_request = append(vw_request, vw_r)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return vw_request, nil
}

func (r *TicketRequestRepo) Created(ctx context.Context, new_ticket *business.TicketRequest) error {
	query := `
		INSERT INTO ticket_request(
			id,
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
			expired_in)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)`

	_, err := r.DB.ExecContext(ctx, query,
		new_ticket.ID,
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
		log.Printf("Error al correr el query")
		return err
	}
	return nil
}
