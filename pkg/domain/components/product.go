package components

import "context"

type Product struct {
	ID         string  `json:"id" db:"id"`
	NAME       string  `json:"name" db:"name"`
	PRODUCT_ID int     `json:"product_id" db:"product_id"`
	DESC       *string `json:"description" db:"description"`
	IS_ACTIVE  bool    `json:"is_active" db:"is_active"`
	CREATEDAT  *string `json:"created_at" db:"created_at"`
}

type ValidateProduct struct {
	NAME       string `json:"id" validate:"required, max=200"`
	PRODUCT_ID int    `json:"product_id" validate:"required"`
}

type ProductRepo interface {
	GetAll(ctx context.Context) ([]*Product, error)
	Created(ctx context.Context, product *Product) error
}
