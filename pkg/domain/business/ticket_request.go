package business

import "context"

type TicketRequest struct {
	ID              string `json:"id" db:"id"`
	REPORTED_BY     string `json:"reported_by_customer_id" db:"reported_by_customer_id"`
	AFFECTED_C      string `json:"affected_customer_id" db:"affected_customer_id"`
	TICKET_BCV      string `json:"ticket_bcv" db:"ticket_bcv"`
	ENVIROMENT      string `json:"enviroment" db:"enviroment"`
	PRIORITY        string `json:"priority" db:"priority"`
	TYPE_REQUEST    string `json:"type_request" db:"type_request"`
	SLA_ID          int    `json:"sla_id" db:"sla_id"`
	COMPONENT_ID    int    `json:"components_id" db:"components_id"`
	SUBCOMPONENT_ID int    `json:"subcomponents_id" db:"subcomponents_id"`
	CONTACT_ID      int    `json:"contact_id" db:"contact_id"`
	TOPIC           string `json:"topic" db:"topic"`
	DESCRIPTION     string `json:"description" db:"description"`
	CREATED_BY      string `json:"created_by" db:"created_by"`
	CREATED_AT      string `json:"created_at" db:"created_at"`
	EXPIRED_IN      string `json:"expired_in" db:"expired_in"`
}

type ValidateTicketRequest struct {
	ID              string `json:"id" binding:"required, min=13,max=13"`
	REPORTED_BY     string `json:"reported_by_customer_id" binding:"required, min=6,max=15"`
	AFFECTED_C      string `json:"affected_customer_id" binding:"required, min=6,max=20"`
	TICKET_BCV      string `json:"ticket_bcv" binding:"required, min=6,max=6"`
	ENVIROMENT      string `json:"enviroment" binding:"required, max=100"`
	PRIORITY        string `json:"priority" binding:"required, max=100"`
	TYPE_REQUEST    string `json:"type_request" binding:"required, max=100"`
	SLA_ID          int    `json:"sla_id" binding:"required"`
	COMPONENT_ID    int    `json:"components_id" binding:"required"`
	SUBCOMPONENT_ID int    `json:"subcomponents_id" binding:"required"`
	CONTACT_ID      int    `json:"contact_id" binding:"required"`
	TOPIC           string `json:"topic" binding:"required"`
	DESCRIPTION     string `json:"description" binding:"required"`
	CREATED_BY      string `json:"created_by" binding:"required, min=6,max=15"`
	CREATED_AT      string `json:"created_at" binding:"required"`
	EXPIRED_IN      string `json:"expired_in" binding:"required"`
}

type InterfaceTicketRequest interface {
	GetAll(ctx context.Context) ([]*TicketRequest, error)
	Created(ctx context.Context, new_ticket *TicketRequest) error
}
