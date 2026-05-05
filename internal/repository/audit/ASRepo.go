package audit

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sydesk/pkg/domain/audit"

	"github.com/google/uuid"
)

type asRepo struct {
	DB *sql.DB
}

func NewAsRepo(db *sql.DB) audit.AuditServInterface {
	return &asRepo{DB: db}
}

func (r *asRepo) GetAll(ctx context.Context, cpr_id uuid.UUID) ([]*audit.AuditServ, error) {

	query := `
		SELECT 
			id,
			cpr_id,
			start_at,
			environment,
			services,
			description,
			end_at,
			last_operation,
			created_at
		FROM audit_status WHERE cpr_id = $1 ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query, cpr_id)

	if err != nil {
		log.Printf("error al correr el query context %v", err.Error())
		return nil, err
	}
	defer rows.Close()

	audit_status := make([]*audit.AuditServ, 0)

	for rows.Next() {
		asRows := &audit.AuditServ{}
		// variable temporar para recibien el json
		var servicesRaw []byte

		err := rows.Scan(
			&asRows.ID,
			&asRows.CPR_ID,
			&asRows.START_AT,
			&asRows.ENVIRONMENT,
			&servicesRaw,
			&asRows.DESCRIPTION,
			&asRows.END_AT,
			&asRows.LAST_OPERATION,
			&asRows.CREATED_AT,
		)

		if err != nil {
			return nil, err
		}

		if len(servicesRaw) > 0 {
			if err := json.Unmarshal(servicesRaw, &asRows.SERVICES); err != nil {
				log.Printf("error al deserealizar el json")
				return nil, err
			}
		}

		audit_status = append(audit_status, asRows)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return audit_status, nil
}

func (r *asRepo) Created(ctx context.Context, audit *audit.AuditServ) error {

	// conversion del slider en json
	servicesJSON, errJson := json.Marshal(audit.SERVICES)
	if errJson != nil {
		return fmt.Errorf("error marshalling services %w", errJson)
	}

	query := `
		INSERT INTO audit_status
			(
				cpr_id,
				start_at,
				environment,
				services,
				activities,
				description,
				end_at,
				last_operation
			)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8)
	`
	_, err := r.DB.ExecContext(ctx, query,
		audit.CPR_ID,
		audit.START_AT,
		audit.ENVIRONMENT,
		servicesJSON,
		audit.ACTIVITIES,
		audit.DESCRIPTION,
		audit.END_AT,
		audit.LAST_OPERATION,
	)

	if err != nil {
		log.Printf("error al correr el insert %v", err.Error())
		return err
	}
	return nil

}
