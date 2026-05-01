package business

import "context"

type Customer struct {
	ID            int     `json:"id" db:"id"`
	RIF           string  `json:"rif" db:"rif"`
	NAME          string  `json:"name" db:"name"`
	CHANNEL       string  `json:"channel" db:"channel"`
	WS_GROUP      string  `json:"ws_group" db:"ws_group"`
	UID_SYPAGO    *string `json:"uid_sypago" db:"uid_sypago"`
	CREATEDAT     string  `json:"created_at" db:"created_at"`
	PRODUCT_ID    int     `json:"product_id" db:"product_id"`
	PRODUCT_ROLE  int     `json:"product_role_id" db:"product_role_id"`
	NOMBRE_BANCO  string  `json:"nombre_banco" db:"nombre_banco"`
	CONTACT_NAME  string  `json:"contact_name" db:"contact_name"`
	CONTACT_LAST  string  `json:"contact_surname" db:"contact_surname"`
	CONTACT_EMAIL string  `json:"contact_email" db:"contact_email"`
	CONTACT_PHONE string  `json:"contact_phone" db:"contact_phone"`
}

type ValidateCustomer struct {
	RIF        string `json:"rif" binding:"required, max=15"`
	NAME       string `json:"name" binding:"required, max=255.min=2"`
	CHANNEL    string `json:"channel" binding:"required, max=100, min=3"`
	WS_GROUP   string `json:"ws_group" binding:"required, min=2, max=100"`
	UID_SYPAGO string `json:"uid_sypago" binding:"required, max=200"`
}

type CustomerRepo interface {
	GetAll(ctx context.Context) ([]*Customer, error)
	Created(ctx context.Context, customer *Customer) error

	// --- MODULO AUDITORIA ---//
	GetAuditCustomer(ctx context.Context, customer_id int) ([]*Customer, error)
}
