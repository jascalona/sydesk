package audit

import "context"

type ItemActivities struct {
	ID         int    `json:"id" db:"od"`
	PRODCUT_ID int    `json:"product_id" db:"product_id"`
	NAME       string `json:"name" db:"name"`
	CREATED_AT string `json:"created_at" db:"created_at"`
}

type ValidationItem struct {
	PRODUCT_ID *int    `json:"product_id" binding:"required"`
	NAME       *string `json:"name" binding:"required,max=255"`
}

type InterfaceItem interface {
	GetAll(ctx context.Context) ([]*ItemActivities, error)
	Created(ctx context.Context, item *ItemActivities) error
}
