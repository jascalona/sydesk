package components

import "context"

type Components struct {
	ID         *int    `json:"id" db:"id"`
	NAME       string  `json:"name" db:"name"`
	PRODUCT_ID int     `json:"product_id" db:"product_id"`
	DESC       string  `json:"description" db:"description"`
	IS_ACTIVE  bool    `json:"is_active" db:"is_active"`
	CREATEDAT  *string `json:"created_at" db:"created_at"`
}

type ValidationComponents struct {
	NAME       string `json:name validate:"required, max=200"`
	PRODUCT_ID int    `json:product_id validate:"required"`
}

type ComponentsRepo interface {
	GetAll(ctx context.Context) ([]*Components, error)
	Created(ctx context.Context, components *Components) error
}
