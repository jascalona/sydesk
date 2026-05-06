package components

import "context"

type Priority struct {
	ID        int    `json:"id" db:"id"`
	NAME      string `json:"name" db:"name"`
	CREATEDAT string `json:"created_at" db:"created_at"`
}

type ValidatePriority struct {
	NAME string `json:"name" validate="required,max=100"`
}

type PriorityRepo interface {
	GetAll(ctx context.Context) ([]*Priority, error)
}
