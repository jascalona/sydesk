package components

import "context"

type Departament struct {
	ID         int    `json:"id" db:"id"`
	NAME       string `json:"name" db:"name"`
	IS_ACTIVE  bool   `json:"is_active" db:"is_active"`
	CREATED_AT string `json:"created_at" db:"created_at"`
}

type ValidateDepartament struct {
	NAME string `json:"name" binding:"required"`
}

type InterfaceDepartament interface {
	GetAll(ctx context.Context) ([]*Departament, error)
}
