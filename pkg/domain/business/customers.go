package business

import (
	"context"
)

type Customer struct {
	ID         int     `json:"id" db:"id"`
	RIF        string  `json:"rif" db:"rif"`
	NAME       string  `json:"name" db:"name"`
	CHANNEL    *string `json:"channel" db:"channel"`
	WS_GROUP   *string `json:"ws_group" db:"ws_group"`
	UID_SYPAGO *string `json:"uid_sypago" db:"uid_sypago"`
	CREATEDAT  string  `json:"created_at" db:"created_at"`
}

type ValidateCustomer struct {
	RIF        string `json:"rif" binding:"required,max=15"`
	NAME       string `json:"name" binding:"required,max=255,min=2"`
	CHANNEL    string `json:"channel" binding:"max=100"`
	WS_GROUP   string `json:"ws_group" binding:"max=100"`
	UID_SYPAGO string `json:"uid_sypago" binding:"max=200"`
}

type CustomerRepo interface {
	GetAll(ctx context.Context) ([]*Customer, error)
	Created(ctx context.Context, customer *Customer) error
}
