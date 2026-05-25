package business

import "context"

type Contact struct {
	ID          int    `json:"id" db:"id"`
	CUSTOMER_ID string `json:"customer_id" db:"customer_id"`
	NAME        string `json:"name" db:"name"`
	SURNAME     string `json:"surname" db:"surname"`
	EMAIL       string `json:"email" db:"email"`
	PHONE       string `json:"phone" db:"phone"`
	CREATED_AT  string `json:"created_at" db:"created_at"`
}

type ContactValidation struct {
	CUSTOMER_ID string `json:"customer_id" binding:"required,max=15,min=6"`
	NAME        string `json:"name" binding:"required,max=225"`
	SURNAME     string `json:"surname" binding:"required,max=225"`
	EMAIL       string `json:"email" binding:"required,email"`
	PHONE       string `json:"phone" binding:"required,max=11,min=11"`
}

type InterfaceContact interface {
	GetAll(ctx context.Context) ([]*Contact, error)
	Created(ctx context.Context, contact *Contact) error
}
