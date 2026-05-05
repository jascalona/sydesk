package audit

import (
	"context"

	"github.com/google/uuid"
)

type Audit struct {
	ID            int       `json:"id" db:"id"`
	RIF           string    `json:"rif" db:"rif"`
	NAME          string    `json:"name" db:"name"`
	CHANNEL       *string   `json:"channel" db:"channel"`
	WS_GROUP      *string   `json:"ws_group" db:"ws_group"`
	UID_SYPAGO    *string   `json:"uid_sypago" db:"uid_sypago"`
	CREATEDAT     string    `json:"created_at" db:"created_at"`
	PRODUCT_ID    int       `json:"product_id" db:"product_id"`
	PRODUCT_ROLE  uuid.UUID `json:"product_role_id" db:"id"`
	NOMBRE_BANCO  *string   `json:"name_ibp" db:"name"`
	CONTACT_NAME  *string   `json:"contact_name" db:"name"`
	CONTACT_LAST  *string   `json:"contact_surname" db:"urname"`
	CONTACT_EMAIL *string   `json:"contact_email" db:"email"`
	CONTACT_PHONE *string   `json:"contact_phone" db:"phone"`
}

type AuditCustomer interface {
	GetAuditCustomer(ctx context.Context, parentID uuid.UUID) ([]*Audit, error)
}
