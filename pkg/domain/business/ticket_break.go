package business

import "context"

type TicketBreak struct {
	ID         int     `json:"id" db:"id"`
	TICKET_ID  string  `json:"ticket_id" db:"ticket_id"`
	STS_P      string  `json:"status_paused" db:"status_paused"`
	START_AT   *string `json:"start_at" db:"start_at"`
	END_AT     *string `json:"end_at" db:"end_at"`
	CREATED_AT *string `json:"created_at" db:"created_at"`
}

type ValidateTicketBreak struct {
	TICKET_ID string `json:"ticket_id" binding:"required,min=13,max=13"`
	STS_P     string `json:"status_paused" binding:"required,max=50"`
	START_AT  string `json:"start_at" binding:"omitempty,datetime=2006-01-02 15:04:05"`
	END_AT    string `json:"end_at" binding:"omitempty,datetime=2006-01-02 15:04:05"`
}

type InterfaceBreak interface {
	GetAll(ctx context.Context) ([]*TicketBreak, error)
	Created(ctx context.Context, t_break *TicketBreak) error
	GetTicketId(ctx context.Context, ticket_id string) ([]*TicketBreak, error)
}
