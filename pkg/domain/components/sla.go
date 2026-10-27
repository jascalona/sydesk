package components

import "context"

type SLA struct {
	ID           int    `json:"id" db:"id"`
	NAME         string `json:"name" db:"name"`
	UNIT_TIME_ID int    `json:"unit_time_id" db:"unit_time_id"`
	AMOUNT       int    `json:"amount" db:"amount"`
	CREATED_AT   string `json:"created_at" db:"created_at"`
}

type ValidateSLA struct {
	NAME         string `json:"name" binding:"required,max=255"`
	UNIT_TIME_ID int    `json:"unit_time_id" binding:"required"`
	AMOUNT       int    `json:"amount" binding:"required"`
}

type InterfaceSLA interface {
	GetAll(ctx context.Context) ([]*SLA, error)
}
