package audit

import (
	"context"

	"github.com/google/uuid"
)

type AuditServ struct {
	ID             uuid.UUID `json:"id" db:"id"`
	CPR_ID         uuid.UUID `json:"cpr_id" db:"cpr_id"`
	START_AT       *string   `json:"start_at" db:"start_at"`
	ENVIRONMENT    string    `json:"environment" db:"environment"`
	SERVICES       *[]string `json:"services" db:"services"`
	DESCRIPTION    string    `json:"description" db:"description"`
	ACTIVITIES     string    `json:"activities" db:"activities"`
	END_AT         *string   `json:"end_at" db:"end_at"`
	LAST_OPERATION *string   `json:"last_operation" db:"last_operation"`
	CREATED_AT     string    `json:"created_at" db:"created_at"`
}

type ValidateAS struct {
	CPR_ID         uuid.UUID `json:"cpr_id" binding:"required"`
	START_AT       string    `json:"start_at" binding:"omitempty,datetime=2006-01-02 15:04:05"`
	ENVIRONMENT    string    `json:"environment" binding:"required"`
	SERVICES       []string  `json:"services" binding:"required,dive,required"`
	ACTIVITIES     string    `json:"activities" binding:"required,max=225"`
	END_AT         string    `json:"end_at" binding:"omitempty,datetime=2006-01-02 15:04:05"`
	LAST_OPERATION string    `json:"last_operation" binding:"omitempty,datetime=2006-01-02 15:04:05"`
	DESCRIPTION    string    `json:"description" binding:"required"`
}

type AuditServInterface interface {
	GetAll(ctx context.Context, cpr_id uuid.UUID) ([]*AuditServ, error)
	Created(ctx context.Context, audit *AuditServ) error
}
