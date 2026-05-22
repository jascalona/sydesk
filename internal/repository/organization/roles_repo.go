package organization

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain"
)

type rolesRepo struct {
	DB *sql.DB
}

func NewRoleRepo(db *sql.DB) domain.RolesRepo {
	return &rolesRepo{DB: db}
}

func (r *rolesRepo) GetAll(ctx context.Context) ([]*domain.Roles, error) {
	query := `SELECT id, name, created_at FROM roles`

	rows, err := r.DB.QueryContext(ctx, query)

	if err != nil {
		log.Printf("Error al correr el query context")
		return nil, err
	}

	roles := make([]*domain.Roles, 0)

	for rows.Next() {
		rol := &domain.Roles{}
		err := rows.Scan(
			&rol.ID,
			&rol.NAME,
			&rol.CREATEDAT,
		)
		if err != nil {
			return nil, err
		}

		roles = append(roles, rol)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *rolesRepo) Created(ctx context.Context, rol *domain.Roles) error {

	query := `INSERT INTO roles(name)VALUES($1)`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		rol.NAME,
	)

	if err != nil {
		log.Printf("error al correr el query", err.Error())
		return err
	}

	return nil
}
