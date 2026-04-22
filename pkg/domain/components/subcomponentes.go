package components

import "context"

type Subcomponents struct {
	ID           int     `json:"id" db:"id"`
	NAME         string  `json:"name" db:"name"`
	COMPONENT_ID int     `json:"component_id" db:"component_id"`
	DESC         *string `json:"description" db:"description"`
	IS_ACTIVE    bool    `json:"is_active" db:"is_active"`
	CREATEDAT    string  `json:"created_at" db:"created_at"`
}

type ValidateSubcomponents struct {
	NAME         string `json:name validate:"required, max=200"`
	COMPONENT_ID int    `json:component_id validate:"required"`
}

type SubcomponentsRepo interface {
	GetAll(ctx context.Context) ([]*Subcomponents, error)
}
