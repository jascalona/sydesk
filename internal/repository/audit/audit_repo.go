package audit

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain/audit"

	"github.com/google/uuid"
)

type auditRepo struct {
	DB *sql.DB
}

func NewAuditRepo(db *sql.DB) audit.AuditCustomer {
	return &auditRepo{DB: db}
}

func (r *auditRepo) GetAuditCustomer(ctx context.Context, parentID uuid.UUID) ([]*audit.Audit, error) {

	query := `
 		SELECT
	   		cli.id,
			cli.rif,
			cli.name,
			cli.channel,
			cli.ws_group,
			cli.uid_sypago,
			cli.created_at,

			cpr.product_id,
			cpr.id AS product_role_id,

			ibp.name AS name_ibp,

			cont.name AS contact_name,
			cont.surname AS contact_surname,
			cont.email AS contact_email,
			cont.phone AS contact_phone

		FROM customer_product_roles cpr
		INNER JOIN customers cli ON cpr.customer_id = cli.id
		LEFT JOIN contact cont ON cont.customer_id = cli.id
		INNER JOIN customer_product_roles prod_padre ON cpr.parent_id = prod_padre.id
		INNER JOIN customers ibp ON prod_padre.customer_id = ibp.id
		WHERE cpr.parent_id = $1
		 -- AND ($2::int = 0 OR cli.id = $2)
		 -- AND cli.id != ibp.id
		 -- AND ibp.name ILIKE '%Bancaribe%';
	`

	rows, err := r.DB.QueryContext(ctx, query, parentID)
	if err != nil {
		log.Printf("Error al correr el query context %v", err.Error())
		log.Println("ESTOY EXPLOTANDO AQUI PAPU")
		return nil, err
	}
	defer rows.Close()

	auditCustom := make([]*audit.Audit, 0)
	for rows.Next() {
		ac := &audit.Audit{}
		err := rows.Scan(
			&ac.ID,
			&ac.RIF,
			&ac.NAME,
			&ac.CHANNEL,
			&ac.WS_GROUP,
			&ac.UID_SYPAGO,
			&ac.CREATEDAT,
			&ac.PRODUCT_ID,
			&ac.PRODUCT_ROLE,
			&ac.NOMBRE_BANCO,
			&ac.CONTACT_NAME,
			&ac.CONTACT_LAST,
			&ac.CONTACT_EMAIL,
			&ac.CONTACT_PHONE,
		)

		if err != nil {
			return nil, err
		}
		auditCustom = append(auditCustom, ac)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return auditCustom, nil
}
