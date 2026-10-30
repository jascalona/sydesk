package business

import "context"

type Notes struct {
	ID int `json:"id" db:"id"`
	//TICKET_ID   *string `json:"ticket_id" db:"ticket_id"` reservar para futuras implementaciones
	TASK_ID     string  `json:"task_id" db:"task_id"`
	TOPIC       string  `json:"topic" db:"topic"`
	DESCRIPTION string  `json:"description" db:"description"`
	CREATED_BY  string  `json:"created_by" db:"created_by"`
	CREATED_AT  *string `json:"created_at" db:"created_at"`
	UPDATE_AT   *string `json:"update_at" db:"update_at"`
}

type ValidateNotes struct {
	TASK_ID     string `json:"task_id" binding:"required,max=6,min=6"`
	TOPIC       string `json:"topic" binding:"required,max=255"`
	DESCRIPTION string `json:"description" binding:"required"`
	CREATED_BY  string `json:"created_by" binding:"required,max=15"`
}

type InterfaceNotes interface {
	GetByTaskId(ctx context.Context, task_id string) ([]*Notes, error)
	CreatedNotes(ctx context.Context, notes *Notes) error
	UpdateNotes(ctx context.Context, note_id *Notes) error
}
