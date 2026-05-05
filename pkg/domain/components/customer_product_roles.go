package components

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

// NullableUUID deserializa JSON null, cadena vacía "" o un UUID válido. Vacío → UUID nil (sin error).
type NullableUUID struct {
	UUID *uuid.UUID
}

func (n *NullableUUID) UnmarshalJSON(data []byte) error {
	if n == nil {
		return nil
	}
	data = bytes.TrimSpace(data)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		n.UUID = nil
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if strings.TrimSpace(s) == "" {
		n.UUID = nil
		return nil
	}
	id, err := uuid.Parse(s)
	if err != nil {
		return err
	}
	n.UUID = &id
	return nil
}

type CustomerProductRole struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	CUSTOMER_ID int        `json:"customer_id" db:"customer_id"`
	PRODUCT_ID  int        `json:"product_id" db:"product_id"`
	ROLE_ID     int        `json:"role_id" db:"role_id"`
	PARENT_ID   *uuid.UUID `json:"parent_id" db:"parent_id"`
	CREATEDAT   string     `json:"created_at" db:"created_at"`
}

type ValidationCustomerProductRole struct {
	CUSTOMER_ID int          `json:"customer_id" binding:"required"`
	PRODUCT_ID  int          `json:"product_id" binding:"required"`
	ROLE_ID     int          `json:"role_id" binding:"required"`
	PARENT_ID   NullableUUID `json:"parent_id"`
}

type CustomerProductRoleRepo interface {
	GetAll(ctx context.Context) ([]*CustomerProductRole, error)
	Created(ctx context.Context, cpr *CustomerProductRole) error
}
