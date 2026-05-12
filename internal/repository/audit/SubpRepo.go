package audit

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sydesk/pkg/domain/audit"
)

type SupRepo struct {
	DB *sql.DB
}

func NewSupRepo(db *sql.DB) audit.InterfaceSup {
	return &SupRepo{DB: db}
}

func (r *SupRepo) GetAll(ctx context.Context) ([]*audit.SupServ, error) {
	query :=
		`	SELECT
				id,
				name,
				product_id,
				created_at
			FROM subproduct ORDER BY created_at DESC		
		`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		log.Printf("error al correr el query context %v", err.Error())
		return nil, err
	}
	defer rows.Close()

	subp := make([]*audit.SupServ, 0)

	for rows.Next() {
		sp := &audit.SupServ{}
		err := rows.Scan(
			&sp.ID,
			&sp.NAME,
			&sp.PRODUCT_ID,
			&sp.CREATED_AT,
		)
		if err != nil {
			return nil, fmt.Errorf("error en el repo")
		}
		subp = append(subp, sp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error %v", err.Error())
	}

	return subp, nil
}
