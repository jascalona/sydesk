package components

import (
	"context"
	"database/sql"
	"sydesk/pkg/domain/components"
)

type enviromentRepo struct {
	DB *sql.DB
}

func NewEnviromentRepo(db *sql.DB) components.EnviromentRepo {
	return &enviromentRepo{DB: db}
}

func (r *enviromentRepo) GetAll(ctx context.Context) ([]*components.Enviroment, error) {

	query := `SELECT id, name, created_at FROM environment`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	enviroment := make([]*components.Enviroment, 0)
	for rows.Next() {
		enviromentRow := &components.Enviroment{}
		err := rows.Scan(
			&enviromentRow.ID,
			&enviromentRow.NAME,
			&enviromentRow.CREATEDAT,
		)
		if err != nil {
			return nil, err
		}
		enviroment = append(enviroment, enviromentRow)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return enviroment, nil

}
