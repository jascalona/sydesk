package domain

import (
	"context"
)

type Roles struct {
	ID        int    `json:"id" db:"id"`
	NAME      string `json:"name" db:"name"`
	CREATEDAT string `json:"created_at" db:"created_at"`
}

type ValidateRoles struct {
	NAME string `json:"name" binding:"required,max=100,min=2"`
}

type RolesRepo interface {
	GetAll(ctx context.Context) ([]*Roles, error)
	Created(ctx context.Context, rol *Roles) error
}
