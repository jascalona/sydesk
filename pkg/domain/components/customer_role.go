package components

import "context"

type CustomerRole struct {
	ID        int    `json:"id" db:"id"`
	NAME      string `json:"name" db:"name"`
	CREATEDAT string `json:"created_at" db:"created_at"`
}

type ValidatorCustomRole struct {
	NAME string `json:name validate:"required, max=200"`
}

type CustomRoleRepo interface {
	GetAll(ctx context.Context) ([]*CustomerRole, error)
}
