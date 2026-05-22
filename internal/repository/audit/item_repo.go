package audit

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain/audit"
)

type ItemRepo struct {
	BD *sql.DB
}

func NewItemRepo(db *sql.DB) audit.InterfaceItem {
	return &ItemRepo{BD: db}
}

func (r *ItemRepo) GetAll(ctx context.Context) ([]*audit.ItemActivities, error) {

	query := `
		SELECT id, product_id, name, created_at FROM item_activities
	`

	rows, err := r.BD.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error al correr el query context %v", err.Error())
		return nil, err
	}

	items := make([]*audit.ItemActivities, 0)

	for rows.Next() {
		ri := &audit.ItemActivities{}
		err := rows.Scan(
			&ri.ID,
			&ri.PRODCUT_ID,
			&ri.NAME,
			&ri.CREATED_AT,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, ri)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil

}

func (r *ItemRepo) Created(ctx context.Context, item *audit.ItemActivities) error {

	query :=
		`
			INSERT INTO item_activities(
				product_id,
				name
			)VALUES($1,$2)
		`
	_, err := r.BD.ExecContext(ctx, query,
		item.PRODCUT_ID,
		item.NAME,
	)
	if err != nil {
		log.Printf("error al hacer el insert %v", err.Error())
		return err
	}
	return nil

}
