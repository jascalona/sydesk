package components

import "context"

type CustomerProductRole struct {
	ID          int    `json:"id" db:"id"`
	CUSTOMER_ID int    `json:"customer_id" db:"customer_id"`
	PRODUCT_ID  int    `json:"product_id" db:"product_id"`
	ROLE_ID     int    `json:"role_id" db:"role_id"`
	CREATEDAT   string `json:"created_at" db:"created_at"`
}

type ValidationCustomerProductRole struct {
	CUSTOMER_ID int `json:customer_id validate:"required"`
	PRODUCT_ID  int `json:product_id validate:"required"`
	ROLE_ID     int `json:role_id validate:"required"`
}

type CustomerProductRoleRepo interface {
	GetAll(ctx context.Context) ([]*CustomerProductRole, error)
}
