package business

import "context"

type Customer struct {
	ID        int    `json:"id" db:"id"`
	RIF       string `json:"rif" db:"rif"`
	NAME      string `json:"name" db:"name"`
	CREATEDAT string `json:"created_at" db:"created_at"`
}

type ValidateCustomer struct {
	RIF  string `json:rif validate:"required, max=15"`
	NAME string `json:name validate:"max=255"`
}

type CustomerRepo interface {
	GetAll(ctx context.Context) ([]*Customer, error)
	Created(ctx context.Context, customer *Customer) error
}
