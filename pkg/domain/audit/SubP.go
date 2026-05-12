package audit

import "context"

type SupServ struct {
	ID         int    `json:"id" db:"id"`
	NAME       string `json:"name" db:"name"`
	PRODUCT_ID int    `json:"product"`
	CREATED_AT string `json:"created_at" db:"created_at"`
}

type ValidateSuP struct {
	NAME       string `json:"name" binding:"required, max=100"`
	PRODUCT_ID int    `json:"product" binding:"required"`
}

type InterfaceSup interface {
	GetAll(ctx context.Context) ([]*SupServ, error)
}
