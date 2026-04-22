package components

import "context"

type Status struct {
	ID        int    `json:"id" db:"id"`
	NAME      string `json:"name" db:"name"`
	CREATEDAT string `json:"createdAt" db:"created_at"`
}

type ValidateStatus struct {
	NAME string `json:name validate:"required, max=200"`
}

type StatusRepo interface {
	GetAll(ctx context.Context) ([]*Status, error)
}
