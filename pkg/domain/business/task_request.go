package business

import "context"

type TaskRequest struct {
	ID            string  `json:"id" db:"id"`
	TICKET_ID     string  `json:"ticket_id" db:"ticket_id"`
	PRIORITY      int     `json:"priority_id" db:"priority_id"`
	STATUS_ID     int     `json:"status_id" db:"status_id"`
	CREATED_BY    string  `json:"created_by" db:"created_by"`
	ASSIGNED_DPT  int     `json:"assigned_departament" db:"assigned_departament"`
	ASSIGNED_USER string  `json:"assigned_user" db:"assigned_user"`
	TOPIC         string  `json:"topic" db:"topic"`
	DESCRIPTION   string  `json:"description" db:"description"`
	CREATED_AT    string  `json:"created_at" db:"created_at"`
	EXPIRED_IN    string  `json:"expired_in" db:"expired_in"`
	CLOSED_IN     *string `json:"closed_in" db:"closed_in"`
}

type VWTaskRequest struct {
	ID            string  `json:"id" db:"id"`
	TICKET_ID     string  `json:"ticket_id" db:"ticket_id"`
	PRIORITY      string  `json:"priority_name" db:"priority_name"`
	STATUS        string  `json:"status_name" db:"status_name"`
	CREATED_BY    string  `json:"created_by" db:"created_by"`
	ASSIGNED_DPT  string  `json:"assigned_departament_name" db:"assigned_departament_name"`
	ASSIGNED_USER string  `json:"assigned_user_name" db:"assigned_user_name"`
	TOPIC         string  `json:"topic" db:"topic"`
	DESCRIPTION   string  `json:"description" db:"description"`
	CREATED_AT    string  `json:"created_at" db:"created_at"`
	EXPIRED_IN    string  `json:"expired_in" db:"expired_in"`
	CLOSED_IN     *string `json:"closed_in" db:"closed_in"`
}

type ValidateRequestTask struct {
	TICKET_ID     string  `json:"ticket_id" binding:"required,max=13,min=13"`
	PRIORITY      int     `json:"priority_id" binding:"required"`
	STATUS_ID     int     `json:"status_id" bindign:"required"`
	CREATED_BY    string  `json:"created_by" bindign:"required,max=15"`
	ASSIGNED_DPT  int     `json:"assigned_departament" bindign:"required"`
	ASSIGNED_USER string  `json:"assigned_user" bindign:"required,max=15"`
	TOPIC         string  `json:"topic" bindign:"required"`
	DESCRIPTION   string  `json:"description" bindign:"required"`
	EXPIRED_IN    string  `json:"expired_in" db:"expired_in"`
	CLOSED_IN     *string `json:"closed_in" bindign:"omitempty"`
}

type InterfaceTaskRequest interface {
	CreatedRequestTask(ctx context.Context, task *TaskRequest) error
	GetVWTaskRequest(ctx context.Context) ([]*VWTaskRequest, error)
	GetVWTaskRequestByTicket(ctx context.Context, ticket_id string) ([]*VWTaskRequest, error)
}
