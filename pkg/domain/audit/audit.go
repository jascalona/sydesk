package audit

import "context"

type Audit struct {
	ID          int    `json:"id" db:"id"`
	UID_SYPAGO  string `json:"uid_sypago" db:"uid_sypago"`
	RIF         string `json:"rif" db:"rif"`
	NAME        string `json:"name" db:"name"`
	Customer_ID int    `json:"customer_id" db:"customer_id"`
	Channel     string `json:"channel" db:"channel"`
	WS_GROUP    string `json:"ws_group" db:"ws_group"`
	CREATED_AT  string `json:"created_at" db:"created_at"`
}

type ValidationAudit struct {
	UID_SYPAGO  string `json:"uid_sypago" binding:"required,max=200"`
	RIF         string `json:"rif" binding:"required,max=15,min=6"`
	NAME        string `json:"name" binding:"required,max=225"`
	Customer_ID int    `json:"customer_id" binding:"required"`
	Channel     string `json:"channel" binding:"required,max=100"`
	WS_GROUP    string `json:"ws_group" binding:"required",min=2`
}

type AuditRepo interface {
	GetAll(ctx context.Context) ([]*Audit, error)
	Created(ctx context.Context, audit *Audit) error
}
